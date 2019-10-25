package elastic

import (
	"fmt"
	. "logweb/ini"
	"logweb/utils"
	"net/url"
	"sync"

	"github.com/henrylee2cn/faygo"

	"github.com/OwnLocal/goes"
)

type elasticRsp struct {
	Error  string
	Replys map[string]goes.Aggregation
}

type elasticHitsRsp struct {
	Error  string
	Replys goes.Hits
}

type ElasticEngine struct {
	lock   sync.Mutex
	islog  bool
	config ElasticConfig

	espool utils.Pool
}

var elastic *ElasticEngine

func (el *ElasticEngine) init() {
	//el.lock.Lock()
	//defer el.lock.Unlock()

	el.config = Config.Elastic

	if el.config.UsePool == 1 {
		el.newEsPool()
	}

	if el.config.ShowLog == 1 {
		el.showLog(true)
	}
}

func (el *ElasticEngine) showLog(show bool) {
	el.islog = show
}

func (el *ElasticEngine) newEsPool() {
	var err error
	el.espool, err = utils.NewTaskPool(&utils.PoolConfig{
		Factory: func() (interface{}, error) {
			return goes.NewClient(el.config.Host, el.config.Port), nil
		},
		Close: func(interface{}) error {
			return nil
		},
	})
	if err != nil {
		faygo.Warningf("newEsPool Create pool error")
	}
}

func (el *ElasticEngine) getIdleClient() *goes.Client {
	esiface, err := el.espool.Get()
	if err != nil {
		return nil
	}
	es, ok := esiface.(*goes.Client)
	if !ok {
		return nil
	}
	return es
}

func (el *ElasticEngine) search(sindex []string, stype []string, query map[string]interface{}) (*elasticRsp, error) {

	es := el.getIdleClient()
	if es == nil {
		return nil, fmt.Errorf("Can not get elastic client")
	}

	extraArgs := make(url.Values, 1)

	searchResults, err := es.Search(query, sindex, stype, extraArgs)
	if err != nil {
		return nil, err
	}

	if el.islog {
		faygo.Debugf("-------------------\n")
		faygo.Debugf(" %v \n", searchResults.Aggregations)
		faygo.Debugf("-------------------\n")
	}

	return &elasticRsp{Error: searchResults.Error, Replys: searchResults.Aggregations}, nil
}

func (el *ElasticEngine) searchHits(sindex []string, stype []string, query map[string]interface{}) (*elasticHitsRsp, error) {
	es := el.getIdleClient()
	if es == nil {
		return nil, fmt.Errorf("Can not get elastic client")
	}

	extraArgs := make(url.Values, 1)

	searchResults, err := es.Search(query, sindex, stype, extraArgs)
	if err != nil {
		return nil, err
	}

	if el.islog {
		faygo.Debugf("-------------------\n")
		faygo.Debugf(" %v \n", searchResults.Hits)
		faygo.Debugf("-------------------\n")
	}

	return &elasticHitsRsp{Error: searchResults.Error, Replys: searchResults.Hits}, nil
}

func Run() {
	elastic = &ElasticEngine{}
	elastic.init()
}

package requests

import (
	"encoding/json"
	"errors"
	"github.com/isarq/nem-sdk-go/model"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MarketInfo struct {
	BTCBCN struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BCN"`
	BTCBLK struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BLK"`
	BTCBTCD struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BTCD"`
	BTCBTM struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BTM"`
	BTCBTS struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BTS"`
	BTCBURST struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BURST"`
	BTCCLAM struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_CLAM"`
	BTCDASH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_DASH"`
	BTCDGB struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_DGB"`
	BTCDOGE struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_DOGE"`
	BTCEMC2 struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_EMC2"`
	BTCFLDC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_FLDC"`
	BTCFLO struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_FLO"`
	BTCGAME struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_GAME"`
	BTCGRC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_GRC"`
	BTCHUC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_HUC"`
	BTCLTC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_LTC"`
	BTCMAID struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_MAID"`
	BTCOMNI struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_OMNI"`
	BTCNAV struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_NAV"`
	BTCNEOS struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_NEOS"`
	BTCNMC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_NMC"`
	BTCNXT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_NXT"`
	BTCPINK struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_PINK"`
	BTCPOT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_POT"`
	BTCPPC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_PPC"`
	BTCRIC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_RIC"`
	BTCSTR struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_STR"`
	BTCSYS struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_SYS"`
	BTCVIA struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_VIA"`
	BTCXVC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_XVC"`
	BTCVRC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_VRC"`
	BTCVTC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_VTC"`
	BTCXBC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_XBC"`
	BTCXCP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_XCP"`
	BTCXEM struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_XEM"`
	BTCXMR struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_XMR"`
	BTCXPM struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_XPM"`
	BTCXRP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_XRP"`
	USDTBTC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_BTC"`
	USDTDASH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_DASH"`
	USDTLTC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_LTC"`
	USDTNXT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_NXT"`
	USDTSTR struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_STR"`
	USDTXMR struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_XMR"`
	USDTXRP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_XRP"`
	XMRBCN struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_BCN"`
	XMRBLK struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_BLK"`
	XMRBTCD struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_BTCD"`
	XMRDASH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_DASH"`
	XMRLTC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_LTC"`
	XMRMAID struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_MAID"`
	XMRNXT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_NXT"`
	BTCETH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_ETH"`
	USDTETH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_ETH"`
	BTCSC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_SC"`
	BTCBCY struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BCY"`
	BTCEXP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_EXP"`
	BTCFCT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_FCT"`
	BTCRADS struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_RADS"`
	BTCAMP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_AMP"`
	BTCDCR struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_DCR"`
	BTCLSK struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_LSK"`
	ETHLSK struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_LSK"`
	BTCLBC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_LBC"`
	BTCSTEEM struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_STEEM"`
	ETHSTEEM struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_STEEM"`
	BTCSBD struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_SBD"`
	BTCETC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_ETC"`
	ETHETC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_ETC"`
	USDTETC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_ETC"`
	BTCREP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_REP"`
	USDTREP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_REP"`
	ETHREP struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_REP"`
	BTCARDR struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_ARDR"`
	BTCZEC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_ZEC"`
	ETHZEC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_ZEC"`
	USDTZEC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_ZEC"`
	XMRZEC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"XMR_ZEC"`
	BTCSTRAT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_STRAT"`
	BTCNXC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_NXC"`
	BTCPASC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_PASC"`
	BTCGNT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_GNT"`
	ETHGNT struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_GNT"`
	BTCGNO struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_GNO"`
	ETHGNO struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_GNO"`
	BTCBCH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_BCH"`
	ETHBCH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_BCH"`
	USDTBCH struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"USDT_BCH"`
	BTCZRX struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_ZRX"`
	ETHZRX struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_ZRX"`
	BTCCVC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_CVC"`
	ETHCVC struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_CVC"`
	BTCOMG struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_OMG"`
	ETHOMG struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_OMG"`
	BTCGAS struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_GAS"`
	ETHGAS struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"ETH_GAS"`
	BTCSTORJ struct {
		ID            int    `json:"id"`
		Last          string `json:"last"`
		LowestAsk     string `json:"lowestAsk"`
		HighestBid    string `json:"highestBid"`
		PercentChange string `json:"percentChange"`
		BaseVolume    string `json:"baseVolume"`
		QuoteVolume   string `json:"quoteVolume"`
		IsFrozen      string `json:"isFrozen"`
		High24Hr      string `json:"high24hr"`
		Low24Hr       string `json:"low24hr"`
	} `json:"BTC_STORJ"`
}

type MarketInfoBtcPrice struct {
	USD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"USD"`
	AUD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"AUD"`
	BRL struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"BRL"`
	CAD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CAD"`
	CHF struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CHF"`
	CLP struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CLP"`
	CNY struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CNY"`
	DKK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"DKK"`
	EUR struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"EUR"`
	GBP struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"GBP"`
	HKD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"HKD"`
	INR struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"INR"`
	ISK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"ISK"`
	JPY struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"JPY"`
	KRW struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"KRW"`
	NZD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"NZD"`
	PLN struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"PLN"`
	RUB struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"RUB"`
	SEK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SEK"`
	SGD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SGD"`
	THB struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"THB"`
	TWD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"TWD"`
}

// Gets market information from Poloniex api
// return {struct} - A MarketInfo struct
func Xem() (MarketInfo, error) {
	c := Client{}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	node := strings.Split(model.MarketInfo, "//")
	node = strings.Split(node[1], "/")
	//port := node[1]
	c.URL.Host = node[0]
	c.URL.Path = "/public"
	c.URL.Scheme = "https"
	params := map[string]string{"command": "returnTicker"}
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return MarketInfo{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return MarketInfo{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return MarketInfo{}, err
	}

	var data MarketInfo
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return MarketInfo{}, err
	}
	return data, nil
}

// Gets BTC price from blockchain.info API
// return {object} - A MarketInfo object
func Btc() (MarketInfoBtcPrice, error) {
	c := Client{}
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	node := strings.Split(model.BtcPrice, "//")
	node = strings.Split(node[1], "/")
	//port := node[1]
	c.URL.Host = node[0]
	c.URL.Path = "/ticker"
	c.URL.Scheme = "https"
	params := map[string]string{"cors": "true"}
	req, err := c.buildReq(params, nil, http.MethodGet)
	if err != nil {
		return MarketInfoBtcPrice{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return MarketInfoBtcPrice{}, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err := errors.New(string(byteArray))
		return MarketInfoBtcPrice{}, err
	}

	var data MarketInfoBtcPrice
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return MarketInfoBtcPrice{}, err
	}
	return data, nil
}

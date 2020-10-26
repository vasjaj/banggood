package client

import (
	"fmt"
	"time"

	"golang.org/x/text/language"
)

var (
	en = language.English.String()
)

func (c client) translateURL(token, productID, poaID, warehouse, currency string) string {
	return fmt.Sprintf("%s/product/Translate?access_token=%s&lang=%s&product_id=%s&poa_id=%s&warehouse=%s&currency=%s", c.BaseURL, token, en, productID, poaID, warehouse, currency)
}

func (c client) getProductPriceURL(token, productID, poaID, warehouse, currency string) string {
	return fmt.Sprintf("%s/product/GetProductPrice?access_token=%s&lang=%s&product_id=%s&poa_id=%s&warehouse=%s&currency=%s", c.BaseURL, token, en, productID, poaID, warehouse, currency)
}

func (c client) getAccessTokenURL() string {
	return fmt.Sprintf("%s/getAccessToken?app_id=%s&app_secret=%s", c.BaseURL, c.AppID, c.AppSecret)
}

func (c client) getCategoryListURL(token string, page *int) string {
	return fmt.Sprintf("%s/category/getCategoryList?access_token=%s&lang=%s&page=%s", c.BaseURL, token, en, page)
}

func (c client) getProductListURL(token, categoryID string, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd *time.Time, page *int) string {
	return fmt.Sprintf("%s/product/getProductList?access_token=%s&lang=%s&cat_id=%sadd_date_start=%s&add_date_end=%s&modify_date_start=%s&modify_date_end=%s&page=%s", c.BaseURL, token, en, categoryID, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd, page)
}

func (c client) getProductInfoURL(token, productID, currency string) string {
	return fmt.Sprintf("%s/product/getProductInfo?access_token=%s&lang=%s&product_id=%s&currency=%s", c.BaseURL, token, en, productID, currency)
}

func (c client) getShipmentsURL(token, productID, warehouse, country, poaID, currency string, quantity int) string {
	return fmt.Sprintf("%s/product/getShipments?access_token=%s&lang=%s&product_id=%s&warehouse=%s&country=%s&poa_id=%s&quantity=%s&currency=%s", c.BaseURL, token, en, productID, warehouse, country, poaID, quantity, currency)
}

func (c client) importOrderURL() string {
	return fmt.Sprintf("%s/importOrder", c.BaseURL)
}

func (c client) getOrderInfoURL(token, saleRecordID string) string {
	return fmt.Sprintf("%s/order/getOrderInfo?access_token=%s&lang=%s&sale_record_id=%s&", c.BaseURL, token, en, saleRecordID)
}

func (c client) getTrackInfoURL(token, orderID string) string {
	return fmt.Sprintf("%s/getTrackInfo?access_token=%s&lang=%s&order_id=%s", c.BaseURL, token, en, orderID)
}

func (c client) getOrderHistoryURL(token, saleRecordID, orderID string) string {
	return fmt.Sprintf("%s/getOrderHistory?access_token=%s&lang=%s&sale_record_id=%s&order_id=%s", c.BaseURL, token, en, saleRecordID, orderID)
}

func (c client) getCountriesURL(token string) string {
	return fmt.Sprintf("%s/common/getCountries?access_token=%s&lang=%s&", c.BaseURL, token, en)
}

func (c client) getStockURL(token, productID string) string {
	return fmt.Sprintf("%s/product/getStocks?access_token=%s&lang=%s&product_id=%s", c.BaseURL, token, en, productID)
}

func (c client) getProductUpdateListURL(token string, minutes, page int) string {
	return fmt.Sprintf("%s/product/getProductUpdateList?access_token=%s&lang=%s&minutes=%s&page=%s", c.BaseURL, token, en, minutes, page)
}

func (c client) getLimitPriceBrandURL(token string, page int) string {
	return fmt.Sprintf("%s/product/getLimitPriceBrand?access_token=%s&page=%s", c.BaseURL, token, page)
}

func (c client) getBrandLimitPriceListURL(token, brandID string, page int) string {
	return fmt.Sprintf("%s/product/getBrandLimitPriceList?access_token=%s&page=%s&brand_id=", c.BaseURL, token, page, brandID)
}

package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultURL     = "https://api.banggood.com"
	defaultTestURL = "https://apibeta.banggood.com&apiTest=1&"
	pageFrom       = 1
)

type BanggoodClient interface {
	Translate(ctx context.Context, token, productID, poaID, warehouse, currency string) (TranslateResponse, error)
	GetProductPrice(ctx context.Context, token, productID, poaID, warehouse, currency string) (GetProductPriceResponse, error)
	GetAccessToken(ctx context.Context) (GetAccessTokenResponse, error)
	GetCategoryList(ctx context.Context, token string, page *int) (GetCategoryListResponse, error)
	GetAllCategories(token string) ([]Category, error)
	GetProductList(ctx context.Context, token, categoryID string, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd *time.Time, page *int) (GetProductListResponse, error)
	GetAllProducts(token, categoryID string, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd *time.Time) ([]Product, error)
	GetProductInfo(ctx context.Context, token, productID string, currency *string) (GetProductInfoResponse, error)
	GetShipments(ctx context.Context, token, productID, warehouse, country, poaID, currency string, quantity int) (GetShipmentsResponse, error)
	ImportOrder(ctx context.Context) (ImportOrderResponse, error)
	GetOrderInfo(ctx context.Context, token, saleRecordID string) (GetOrderInfoResponse, error)
	GetTrackInfo(ctx context.Context, token, orderID string) (GetTrackInfoResponse, error)
	GetOrderHistory(ctx context.Context, token, saleRecordID, orderID string) (GetOrderHistoryResponse, error)
	GetCountries(ctx context.Context, token string) (GetCountriesResponse, error)
	GetStock(ctx context.Context, token, productID string) (GetStockResponse, error)
	GetProductUpdateList(ctx context.Context, token string, minutes, page int) (GetProductUpdateListResponse, error)
	GetLimitPriceBrand(ctx context.Context, token string, page int) (GetLimitPriceBrandResponse, error)
	GetBrandLimitPriceList(ctx context.Context, token, brandID string, page int) (GetBrandLimitPriceListResponse, error)
}

func NewDefaultClient(id, secret string) BanggoodClient {
	return client{
		AppID:      id,
		AppSecret:  secret,
		HTTPClient: http.DefaultClient,
		BaseURL:    defaultURL,
	}
}

type client struct {
	AppID      string
	AppSecret  string
	BaseURL    string
	HTTPClient *http.Client
}

func (c client) do(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}
func (c client) Translate(ctx context.Context, token, productID, poaID, warehouse, currency string) (TranslateResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.translateURL(token, productID, poaID, warehouse, currency), nil)
	if err != nil {
		return TranslateResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return TranslateResponse{}, err
	}
	var data TranslateResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetProductPrice(ctx context.Context, token, productID, poaID, warehouse, currency string) (GetProductPriceResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getProductPriceURL(token, productID, poaID, warehouse, currency), nil)
	if err != nil {
		return GetProductPriceResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetProductPriceResponse{}, err
	}
	var data GetProductPriceResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetAccessToken(ctx context.Context) (GetAccessTokenResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getAccessTokenURL(), nil)
	if err != nil {
		return GetAccessTokenResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetAccessTokenResponse{}, err
	}
	var data GetAccessTokenResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetCategoryList(ctx context.Context, token string, page *int) (GetCategoryListResponse, error) {
	// fmt.Println("URL: ", c.getCategoryListURL(token, page))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getCategoryListURL(token, page), nil)
	if err != nil {
		return GetCategoryListResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetCategoryListResponse{}, err
	}
	var data GetCategoryListResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetAllCategories(token string) ([]Category, error) {
	var categories []Category
	page := pageFrom
	for {
		res, err := c.GetCategoryList(context.Background(), token, &page)
		if err != nil {
			return nil, err
		}
		categories = append(categories, res.CategoryList...)
		if res.PageNumber == res.PageTotal {
			break
		}
		page++
	}
	return categories, nil
}

func (c client) GetProductList(ctx context.Context, token, categoryID string, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd *time.Time, page *int) (GetProductListResponse, error) {
	// fmt.Println("URL: ", c.getProductListURL(token, categoryID, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd, page))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getProductListURL(token, categoryID, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd, page), nil)
	if err != nil {
		return GetProductListResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetProductListResponse{}, err
	}
	var data GetProductListResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetAllProducts(token, categoryID string, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd *time.Time) ([]Product, error) {
	var products []Product
	page := pageFrom
	for {
		res, err := c.GetProductList(context.Background(), token, categoryID, addDateStart, addDateEnd, modifyDateStart, modifyDateEnd, &page)
		if err != nil {
			return nil, err
		}
		products = append(products, res.ProductList...)
		if res.PageNumber == res.PageTotal {
			break
		}
		page++
	}
	return products, nil
}

func (c client) GetProductInfo(ctx context.Context, token, productID string, currency *string) (GetProductInfoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getProductInfoURL(token, productID, currency), nil)
	if err != nil {
		return GetProductInfoResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetProductInfoResponse{}, err
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response: ", string(bytes))
	var data GetProductInfoResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetShipments(ctx context.Context, token, productID, warehouse, country, poaID, currency string, quantity int) (GetShipmentsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getShipmentsURL(token, productID, warehouse, country, poaID, currency, quantity), nil)
	if err != nil {
		return GetShipmentsResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetShipmentsResponse{}, err
	}
	var data GetShipmentsResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) ImportOrder(ctx context.Context) (ImportOrderResponse, error) {
	return ImportOrderResponse{}, errors.New("unimplemented")
}

func (c client) GetOrderInfo(ctx context.Context, token, saleRecordID string) (GetOrderInfoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getOrderInfoURL(token, saleRecordID), nil)
	if err != nil {
		return GetOrderInfoResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetOrderInfoResponse{}, err
	}
	var data GetOrderInfoResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetTrackInfo(ctx context.Context, token, orderID string) (GetTrackInfoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getTrackInfoURL(token, orderID), nil)
	if err != nil {
		return GetTrackInfoResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetTrackInfoResponse{}, err
	}
	var data GetTrackInfoResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetOrderHistory(ctx context.Context, token, saleRecordID, orderID string) (GetOrderHistoryResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getOrderHistoryURL(token, saleRecordID, orderID), nil)
	if err != nil {
		return GetOrderHistoryResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetOrderHistoryResponse{}, err
	}
	var data GetOrderHistoryResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetCountries(ctx context.Context, token string) (GetCountriesResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getCountriesURL(token), nil)
	if err != nil {
		return GetCountriesResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetCountriesResponse{}, err
	}
	var data GetCountriesResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetStock(ctx context.Context, token, productID string) (GetStockResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getStockURL(token, productID), nil)
	if err != nil {
		return GetStockResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetStockResponse{}, err
	}
	var data GetStockResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetProductUpdateList(ctx context.Context, token string, minutes, page int) (GetProductUpdateListResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getProductUpdateListURL(token, minutes, page), nil)
	if err != nil {
		return GetProductUpdateListResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetProductUpdateListResponse{}, err
	}
	var data GetProductUpdateListResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetLimitPriceBrand(ctx context.Context, token string, page int) (GetLimitPriceBrandResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getLimitPriceBrandURL(token, page), nil)
	if err != nil {
		return GetLimitPriceBrandResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetLimitPriceBrandResponse{}, err
	}
	var data GetLimitPriceBrandResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

func (c client) GetBrandLimitPriceList(ctx context.Context, token, brandID string, page int) (GetBrandLimitPriceListResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.getBrandLimitPriceListURL(token, brandID, page), nil)
	if err != nil {
		return GetBrandLimitPriceListResponse{}, err
	}
	res, err := c.do(req)
	if err != nil {
		return GetBrandLimitPriceListResponse{}, err
	}
	var data GetBrandLimitPriceListResponse
	return data, json.NewDecoder(res.Body).Decode(&data)
}

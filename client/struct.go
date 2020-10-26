package client

type Page struct {
	PageNumber int `json:"page"`
	PageTotal  int `json:"page_total"`
	PageSize   int `json:"page_size"`
}

type TranslateResponse struct {
}

type GetProductPriceResponse struct {
	TranslatedText string `json:"TranslatedText"`
	Error          int    `json:"error"`
	ErrorMessage   string `json:"errMsg"`
}

type GetAccessTokenResponse struct {
	Code        int    `json:"code"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Category struct {
	CategoryID   string `json:"cat_id"`
	CategoryName string `json:"cat_name"`
	ParentID     string `json:"parent_id"`
}

type GetCategoryListResponse struct {
	Page

	Code          int        `json:"code"`
	CategoryTotal int        `json:"cat_total"`
	Language      string     `json:"lang"`
	CategoryList  []Category `json:"cat_list"`
}

type Product struct {
	ProductID       string `json:"product_id"`
	CategoryID      int    `json:"cat_id"` //TODO: report error
	ProductName     string `json:"product_name"`
	Image           string `json:"img"`
	MetaDescription string `json:"meta_desc"`
	AddDate         string `json:"add_date"`
	ModifyDate      string `json:"modify_date"`
}

type GetProductListResponse struct {
	Page

	Code         int       `json:"code"`
	ProductTotal int       `json:"product_total"`
	Language     string    `json:"lang"`
	ProductList  []Product `json:"product_list"`
}

type GetProductInfoResponse struct {
	PoaList []struct {
		OptionID     string `json:"option_id"`
		OptionName   string `json:"option_name"`
		OptionValues []struct {
			PoaID         string  `json:"poa_id"`
			PoaName       string  `json:"poa_name"`
			Poa           string  `json:"poa"`
			PoaPrice      float32 `json:"poa_price"`
			SmallImage    string  `json:"small_image"`
			ViewImage     string  `json:"view_image"`
			LargeImage    string  `json:"large_image"`
			ListGridImage string  `json:"list_grid_image"`
		} `json:"option_values"`
	} `json:"poa_list"`
	WarehouseList []struct {
		Warehouse      string  `json:"warehouse"`
		WarehousePrice float32 `json:"warehouse_price"`
	} `json:"warehouse_list"`
	ImageList []struct {
		Home       string `json:"home"`
		ListGrid   string `json:"list_grid"`
		Grid       string `json:"grid"`
		Gallery    string `json:"gallery"`
		View       string `json:"view"`
		OtherItems string `json:"other_items"`
		Large      string `json:"large"`
	} `json:"image_list"`
	Description string `json:"description"`
	Code        int    `json:"code"`
	Language    string `json:"lang"`
	Weight      int    `json:"weight"`
	ProductName string `json:"product_name"`
}

type GetShipmentsResponse struct {
	Code           int    `json:"code"`
	Currency       string `json:"currency"`
	ShipMethodCode string `json:"shipmethodcode"`
	ShipMethodName string `json:"shipmethodname"`
	Shipday        string `json:"shipday"`
	Shipfee        string `json:"shipfee"`
}

type ImportOrderRequest struct {
	AccessToken            string `json:"access_token"`
	SaleRecordID           string `json:"sale_record_id"`
	DeliveryName           string `json:"delivery_name"`
	DeliveryCountry        string `json:"delivery_country"`
	DeliveryState          string `json:"delivery_state"`
	DeliveryCity           string `json:"delivery_city"`
	DeliveryStreetAddress  string `json:"delivery_street_address"`
	DeliveryStreetAddress2 string `json:"delivery_street_address2"`
	DeliveryPostcode       string `json:"delivery_postcode"`
	DeliveryTelephone      string `json:"delivery_telephone"`
	ProductTotal           int    `json:"product_total"`
	ProductList            []struct {
		ProductID      string `json:"product_id"`
		PoaID          string `json:"poa_id"`
		Quantity       string `json:"quantity"`
		ShipmethodCode string `json:"shipmethod_code"`
	} `json:"product_list"`
	Language string `json:"lang"`
	Currency string `json:"currency"`
}

type ImportOrderResponse struct {
	SaleRecordID string `json:"sale_record_id"`
	ProductTotal string `json:"product_total"`
	SuccessTotal string `json:"success_total"`
	FailureTotal string `json:"failure_total"`
	FailureList  []struct {
		ProductID        string `json:"product_id"`
		PoaID            string `json:"poa_id"`
		Warehouse        string `json:"warehouse"`
		Quantity         string `json:"quantity"`
		ShipmethodCode   string `json:"shipmethod_code"`
		ErrorDescription string `json:"error_desc"`
	} `json:"failure_list"`
	Code string `json:"code"`
}

type GetOrderInfoResponse struct {
	Code             int `json:"code"`
	SaleRecordIDList []struct {
		SaleRecordID string `json:"sale_record_id"`
		OrderList    []struct {
			OrderID          string  `json:"order_id"`
			Status           string  `json:"status"`
			TotalAmount      float32 `json:"total_amount"`
			Currency         string  `json:"currency"`
			ShipmethodCode   string  `json:"shipment_method"`
			SubAmount        float32 `json:"sub_amount"`
			DropShipDiscount float32 `json:"ds_discount"`
			Shipfee          float32 `json:"shipfee"`
			ShipInsurance    float32 `json:"ship_insurance"`
			TariffInsurance  float32 `json:"tariff_insurance"`
			ProductList      []struct {
				ProductID string   `json:"product_id"`
				Warehouse []string `json:"warehouse"`
				Quantity  int      `json:"quantity"`
				PoaID     string   `json:"poa_id"`
			} `json:"product_list"`
		} `json:"order_list"`
		UserInfo []struct {
			DeliveryName           string `json:"delivery_name"`
			DeliveryCountry        string `json:"delivery_country"`
			DeliveryState          string `json:"delivery_state"`
			DeliveryCity           string `json:"delivery_city"`
			DeliveryStreetAddress  string `json:"delivery_street_address"`
			DeliveryStreetAddress2 string `json:"delivery_steet_address2"`
		} `json:"user_info"`
	} `json:"sale_record_id_list"`
}

type GetTrackInfoResponse struct {
	TrackInfo []struct {
		Event string `json:"event"`
		Time  string `json:"time"`
	} `json:"track_info"`
	Code string `json:"code"`
}

type GetOrderHistoryResponse struct {
	OrderHistory []struct {
		Status  string `json:"status"`
		DateAdd string `json:"date_add"`
	} `json:"order_history"`
	TrackNumber string `json:"track_number"`
	Code        string `json:"code"`
}

type GetCountriesResponse struct {
	Countries []struct {
		CountryID   int    `json:"country_id"`
		CountryName string `json:"country_name"`
	} `json:"countries"`
	Code int `json:"code"`
}

type GetStockResponse struct {
	Stocks []struct {
		Warehouse  string `json:"warehouse"`
		StocksList []struct {
			PoaID         int    `json:"poa_id"`
			Poa           string `json:"poa"`
			Stock         string `json:"stock"`
			StocksMessage string `json:"stocks_msg"`
		} `json:"stocks_list"`
	} `json:"stocks"`
	Code     string `json:"code"`
	Language string `json:"lang"`
}

type GetProductUpdateListResponse struct {
	Page

	Code              int    `json:"code"`
	ProductTotal      int    `json:"product_total"`
	Language          string `json:"lang"`
	UpdateProductList []struct {
		ProductID  string `json:"product_id"`
		State      int    `json:"state"`
		ModifyDate string `json:"modify_date"`
	} `json:"update_product_list"`
}

type GetBrandLimitPriceListResponse struct {
	Page

	Code         int    `json:"code"`
	ProductTotal int    `json:"product_total"`
	Language     string `json:"lang"`
	ProductList  []struct {
		ProductID  string `json:"product_id"`
		Sku        string `json:"sku"`
		Poa        string `json:"poa"`
		LimitPrice string `json:"limit_price"`
	} `json:"product_list"`
}

type GetLimitPriceBrandResponse struct {
	Page

	Code       int    `json:"code"`
	BrandTotal int    `json:"brand_total"`
	Language   string `json:"lang"`
	BrandList  []struct {
		BrandID string `json:"brand_id"`
		Name    string `json:"name"`
	} `json:"brand_list"`
}

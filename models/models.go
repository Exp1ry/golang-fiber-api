package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NFTBrokerModel struct {
ID primitive.ObjectID `json:"_id"`
Name string `json:"name"`
Description string `json:"description"`
}
type CryptoBrokerModel struct {
ID primitive.ObjectID `json:"_id"`
Name string `json:"name"`
Description string `json:"description"`
}
type StockBrokerModel struct {
ID primitive.ObjectID `json:"_id"`
Name string `json:"name"`
Description string `json:"description"`
}
type ForexBrokerModel struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
    RealID int64 `json:"id" bson:"id"`
    Name                   string    `json:"name" bson:"name"`
    Slug                   string    `json:"slug" bson:"slug,omitempty"`
    Active                 bool      `json:"active" bson:"active,omitempty"`
    HeadquartersAddress    string    `json:"headquartersAddress" bson:"headquartersAddress,omitempty"`
    Email                  string    `json:"email" bson:"email,omitempty"`
    Description            string    `json:"description" bson:"description,omitempty"`
    Country                string    `json:"country" bson:"country,omitempty"`
    PhoneNumber           string    `json:"phoneNumber" bson:"phoneNumber,omitempty"`
    LogoSmall              string    `json:"logoSmall" bson:"logoSmall,omitempty"`
    LogoLarge              string    `json:"logoLarge" bson:"logoLarge,omitempty"`
    SvgSquareLogo          string    `json:"svgSquareLogo" bson:"svgSquareLogo,omitempty"`
    MobileTradingOptions   []string  `json:"mobileTradingOptions" bson:"mobileTradingOptions,omitempty"`
    PubliclyTraded         bool      `json:"publiclyTraded" bson:"publiclyTraded,omitempty"`
    RestrictedCountrie     []string  `json:"restrictedCountrie" bson:"restrictedCountrie,omitempty"`
    NumberOfBonds          int       `json:"numberOfBonds" bson:"numberOfBonds,omitempty"`
    NumberOfCommodities    int       `json:"numberOfCommodities" bson:"numberOfCommodities,omitempty"`
    NumberofCryptocurrency int       `json:"numberofCryptocurrency" bson:"numberofCryptocurrency,omitempty"`
    NumberofEtfs           int       `json:"numberofEtfs" bson:"numberofEtfs,omitempty"`
    NumberofIndices        int       `json:"numberofIndices" bson:"numberofIndices,omitempty"`
    NumberOfStocks         int       `json:"numberOfStocks" bson:"numberOfStocks,omitempty"`
    NumberOfFutures        int       `json:"numberOfFutures" bson:"numberOfFutures,omitempty"`
    NumberOfOptions        int       `json:"numberOfOptions" bson:"numberOfOptions,omitempty"`
    NumberOfCurrencyPairs  int       `json:"numberOfCurrencyPairs" bson:"numberOfCurrencyPairs,omitempty"`
    TradingDeskTypes       []string  `json:"tradingDeskTypes" bson:"tradingDeskTypes,omitempty"`
    TradingPlatforms       []string  `json:"tradingPlatforms" bson:"tradingPlatforms,omitempty"`
    OsTradingPlatforms     []string  `json:"osTradingPlatforms" bson:"osTradingPlatforms,omitempty"`
    CustomerSupportLanguages []string `json:"customerSupportLanguages" bson:"customerSupportLanguages,omitempty"`
    PlatformSupportLanguages []string `json:"platformSupportLanguages" bson:"platformSupportLanguages,omitempty"`
    DepositOptions         []string  `json:"depositOptions" bson:"depositOptions,omitempty"`
    WithdrawlOptions       []string  `json:"withdrawlOptions" bson:"withdrawlOptions,omitempty"`
    SupportedCoins         []string  `json:"supportedCoins" bson:"supportedCoins,omitempty"`
    OfficialWebsiteUrl     string    `json:"officialWebsiteUrl" bson:"officialWebsiteUrl,omitempty"`
    Regulation             []string  `json:"regulation" bson:"regulation,omitempty"`
    RegulationReferenceUrl string    `json:"regulationReferenceUrl" bson:"regulationReferenceUrl,omitempty"`
    Employees              string    `json:"employees" bson:"employees,omitempty"`
    Facebook               string    `json:"facebook" bson:"facebook,omitempty"`
    Twitter                string    `json:"twitter" bson:"twitter,omitempty"`
    FoundationYear         int       `json:"foundationYear" bson:"foundationYear,omitempty"`
   
}
# web-page-scrapper
It contains apis in Golang to scrape web pages and persist them in mysql db.


**steps to use:**
1. Clone the project using `git clone --recursive https://github.com/manoj-JU/web-page-scrapper.git`
2. get inside the repo : `cd web-page-scrapper`
3. start the services using : `docker-compose up`
4. call APIs.
  
# APIs
1.**Scrapes an amazon product page & call another service's api to persist data to mysql database**

**URL:** `http://localhost:8121/scrape/amazon`

**Method:** POST

**Content-Type:** application/json

**Authentication Required:** NO

**Data Examples:**

**Query Params:**

```json
{
   "url":"https://www.amazon.com/Logitech-LIGHTSYNC-Wired-Gaming-Mouse/dp/B07YN82X3B/ref=pd_sim_2/132-8669378-8065121?pd_rd_w=rTwov&pf_rd_p=d3b9006b-884c-4f9d-b9b9-fcc89494e569&pf_rd_r=0ZKHJMPRY70FHWHPWMQD&pd_rd_r=11e7a1c2-7c12-4525-9dad-fb8f76bea8ba&pd_rd_wg=qYDMO&pd_rd_i=B07YN82X3B&psc=1"
}
```

## Success Response:

**Condition:** valid web page provided is valid.

**Code:** `200 OK`


**Content Example:**

```json
{
    "id": 10,
    "title": "Logitech G203 Wired Gaming Mouse, 8,000 DPI, Rainbow Optical Effect LIGHTSYNC RGB, 6 Programmable Buttons, On-Board Memory, Screen Mapping, PC/Mac Computer and Laptop Compatible - Black",
    "image_url": "https://images-na.ssl-images-amazon.com/images/I/61UxfXTUyvL.__AC_SX300_SY300_QL70_ML2_.jpg",
    "description": "World’s No.1 Best Selling Gaming Gear Brand - Based on independent aggregated sales data (FEB ‘19 - FEB’20) of Gaming Keyboard, Mice, & PC Headset in units from: US, CA, CN, JP, KR, TW, TH, ID, DE, FR, RU, UK, SE, TR8, 000 DPI gaming-grade sensor responds precisely to movements. Customize your sensitivity settings to suit the sensitivity you like with Logitech G HUB gaming software and cycle easily through up to 5 DPI settings.Play in color with our most vibrant Lightsync RGB featuring color wave effects customizable across -16.8 million colors. Install Logitech G HUB software to choose from preset colors and animations or make your own. Game-driven, audio visualization and screen mapping options are also available.Play comfortably and with total control. The classic and simple 6-button layout and classic gaming shape is a comfortable time-tested and loved design. Each button can be customized using Logitech G HUB software to simplify tasks.Primary buttons are mechanical and tensioned with durable metal springs for reliability, performance and excellent feel. The crisp clicks and precise feedback delivers a great precision feel to maximize your fun in game.",
    "price": "$29.54",
    "total_reviews": "6049",
    "created_at": "2021-05-19T18:17:30Z",
    "updated_at": "2021-05-19T18:22:13.21Z"
}
```

## Error Response:

**Condition:** If Invalid web page provider.

**Code:** `400 BAD Request`

**Content Example:**

```json:
{
    "message": "Invalid webPage"
}
```

2.**Get all amazon products stored in the database**

**URL:** `http://localhost:8621/amazon/products`

**Method:** GET

**Content-Type:** application/json

**Authentication Required:** NO

**Data Examples:**

## Success Response:

**Condition:** valid web page provided is valid.

**Code:** `200 OK`


**Content Example:**

```json
[
    {
        "id": 9,
        "title": "PlayStation 4 Pro 1TB Console",
        "image_url": "https://images-na.ssl-images-amazon.com/images/I/6118ctEjpoL.__AC_SX300_SY300_QL70_ML2_.jpg",
        "description": "Heighten your experiences. Enrich your adventures. Let the super charged PS4 Pro lead the way4K TV Gaming : PS4 Pro outputs gameplay to your 4K TVMore HD Power: Turn on Boost Mode to give PS4 games access to the increased power of PS4 ProHDR Technology : With an HDR TV, compatible PS4 games display an unbelievably vibrant and life like range of colors",
        "price": "",
        "total_reviews": "11067",
        "created_at": "2021-05-19T16:58:52Z",
        "updated_at": "2021-05-19T18:20:35Z"
    },
    {
        "id": 10,
        "title": "Logitech G203 Wired Gaming Mouse, 8,000 DPI, Rainbow Optical Effect LIGHTSYNC RGB, 6 Programmable Buttons, On-Board Memory, Screen Mapping, PC/Mac Computer and Laptop Compatible - Black",
        "image_url": "https://images-na.ssl-images-amazon.com/images/I/61UxfXTUyvL.__AC_SX300_SY300_QL70_ML2_.jpg",
        "description": "World’s No.1 Best Selling Gaming Gear Brand - Based on independent aggregated sales data (FEB ‘19 - FEB’20) of Gaming Keyboard, Mice, & PC Headset in units from: US, CA, CN, JP, KR, TW, TH, ID, DE, FR, RU, UK, SE, TR8, 000 DPI gaming-grade sensor responds precisely to movements. Customize your sensitivity settings to suit the sensitivity you like with Logitech G HUB gaming software and cycle easily through up to 5 DPI settings.Play in color with our most vibrant Lightsync RGB featuring color wave effects customizable across -16.8 million colors. Install Logitech G HUB software to choose from preset colors and animations or make your own. Game-driven, audio visualization and screen mapping options are also available.Play comfortably and with total control. The classic and simple 6-button layout and classic gaming shape is a comfortable time-tested and loved design. Each button can be customized using Logitech G HUB software to simplify tasks.Primary buttons are mechanical and tensioned with durable metal springs for reliability, performance and excellent feel. The crisp clicks and precise feedback delivers a great precision feel to maximize your fun in game.",
        "price": "$29.54",
        "total_reviews": "6049",
        "created_at": "2021-05-19T18:17:30Z",
        "updated_at": "2021-05-19T18:22:13Z"
    }
]
```

3.**Save product to mysql database**

**URL:** `http://localhost:8621/amazon/products`

**Method:** POST

**Content-Type:** application/json

**Authentication Required:** NO

**Data Examples:**

**Query Params:**

```json
{
        "title": "Logitech G203 Wired Gaming Mouse, 8,000 DPI, Rainbow Optical Effect LIGHTSYNC RGB, 6 Programmable Buttons, On-Board Memory, Screen Mapping, PC/Mac Computer and Laptop Compatible - Black",
        "image_url": "https://images-na.ssl-images-amazon.com/images/I/61UxfXTUyvL.__AC_SX300_SY300_QL70_ML2_.jpg",
        "description": "World’s No.1 Best Selling Gaming Gear Brand - Based on independent aggregated sales data (FEB ‘19 - FEB’20) of Gaming Keyboard, Mice, & PC Headset in units from: US, CA, CN, JP, KR, TW, TH, ID, DE, FR, RU, UK, SE.",
        "price": "$29.54",
        "total_reviews": "6049"
    }
```

## Success Response:

**Condition:** valid web page provided is valid.

**Code:** `200 OK`


**Content Example:**

```json
{
    "ID": 7,
    "CreatedAt": "2021-05-19T18:58:06Z",
    "UpdatedAt": "2021-05-19T19:02:44.203Z",
    "DeletedAt": null,
    "Title": "Logitech G203 Wired Gaming Mouse, 8,000 DPI, Rainbow Optical Effect LIGHTSYNC RGB, 6 Programmable Buttons, On-Board Memory, Screen Mapping, PC/Mac Computer and Laptop Compatible - Black",
    "ImageURL": "https://images-na.ssl-images-amazon.com/images/I/61UxfXTUyvL.__AC_SX300_SY300_QL70_ML2_.jpg",
    "Description": "World’s No.1 Best Selling Gaming Gear Brand - Based on independent aggregated sales data (FEB ‘19 - FEB’20) of Gaming Keyboard, Mice, & PC Headset in units from: US, CA, CN, JP, KR, TW, TH, ID, DE, FR, RU, UK, SE.",
    "Price": "$29.54",
    "TotalReviews": "6049"
}
```

## Error Response:

**Condition:** If Invalid web page provider.

**Code:** `500 Internal Server Error`

**Content Example:**

```json:
{
    "message": "Error"
}
```

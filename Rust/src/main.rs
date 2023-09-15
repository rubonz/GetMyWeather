extern crate serde;
extern crate serde_json;
#[macro_use]
extern crate serde_derive;

use std::io;
extern crate dotenv;

use dotenv::dotenv;
use std::env;

#[derive(Debug, Serialize, Deserialize)]
struct MainData {
    temp: f64,
    feels_like: f64,
}

#[derive(Debug, Serialize, Deserialize)]
struct WeatherData {
    main: MainData,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {

    dotenv().ok();

    println!("Enter your town: ");
    let mut city = String::new();
    io::stdin().read_line(&mut city)?;

    let api_key = env::var("OPENWEATHER_API_KEY")
        .expect("API key not set. Please set the OPENWEATHER_API_KEY environment variable.");

    let url = format!("https://api.openweathermap.org/data/2.5/weather?q={}&units=imperial&lang=en&appid={}", city.trim(), api_key);

    let response = ureq::get(&url).call()?;

    if response.status() < 200 || response.status() >= 300 {
        eprintln!("Error: {}", response.status());
        return Ok(());
    }

    let body = response.into_string()?;
    let weather_data: WeatherData = serde_json::from_str(&body)?;

    let temperature = weather_data.main.temp;
    let temperature_feels = weather_data.main.feels_like;
    println!("Now in the town {}: {:.0}°F", city.trim(), temperature);
    println!("Feels like {:.0}°F", temperature_feels);

    Ok(())
}
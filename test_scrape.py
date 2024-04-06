import requests
from bs4 import BeautifulSoup


def scrape_player_data(player_url):
    response = requests.get(player_url, timeout=10)
    soup = BeautifulSoup(response.text, "html.parser")

    # Extract player details as before
    # Assuming you have correct selectors
    player_name_wrapper = soup.find("div", class_="title")  # Placeholder
    player_name = player_name_wrapper.find("h1").text  # Placeholder

    print(player_name)  # Or insert into database


# Function to scrape the listing page and get player URLs
def scrape_listing_page(listing_url):
    response = requests.get(listing_url, timeout=10)
    soup = BeautifulSoup(response.text, "html.parser")

    # Find all links to player pages - adjust selector as needed
    # Assuming player_links_wrapper is correctly finding the <div> containing player information
    player_links_wrapper = soup.find("div", class_="players")  # This finds the <div>

    # Instead of iterating over player_links_wrapper directly,
    # you should find all <ul> or <li> elements (or whatever contains the player links) inside it
    player_items = player_links_wrapper.find_all(
        "ul"
    )  # Assuming these are the containers

    for player_item in player_items:
        # Now, player_item is a BeautifulSoup object and you can use .find() with keyword arguments on it
        player = player_item.find("li", class_="player")
        name_span = player.find("span", class_="name")
        link = name_span.find("a")
        url = link.get("href")
        scrape_player_data("https://fminside.net" + url)


# Main URL to start scraping from - adjust as needed
listing_url = "https://fminside.net/players"
scrape_listing_page(listing_url)

import requests
from bs4 import BeautifulSoup

file = open("data.csv", "w")


def scrape_player_data(player_url, file):
    response = requests.get(player_url, timeout=10)
    soup = BeautifulSoup(response.text, "html.parser")

    # Extract player details as before
    # Assuming you have correct selectors
    player_wrapper = soup.find("div", id="player")
    player_name_wrapper = soup.find("div", class_="title")  # Placeholder
    player_name = player_name_wrapper.find("h1").text
    print(player_name)
    player_image_style = player_wrapper.find("span").get("style")

    start = (
        player_image_style.find("url(") + 4
    )  # Find the start of the URL and adjust to skip 'url('
    end = player_image_style.find(
        ")", start
    )  # Find the end of the URL, starting from the end of 'url('
    url = player_image_style[start + 2: end]

    file.write(f"Name: {player_name} - ")
    basic_info_wrapper = (
        player_name_wrapper.find(
            "div", class_="meta").find("ul").find_all("li")
    )

    club = basic_info_wrapper[0].find("a").find("span", class_="value").text
    country = basic_info_wrapper[1].find("span").find("a").text

    file.write(f"Club: {club} - Country: {country} - Image URL: {url} - ")

    player_info_wrapper = (
        player_wrapper.find("div", class_="column").find("ul").find_all("li")
    )

    for element in player_info_wrapper:
        field = element.find("span", class_="key").text
        if field == "Caps / Goals" or field == "Unique ID":
            continue
        value = element.find("span", class_="value").text.replace(",", "/")
        file.write(f"{field}: {value} - ")

    attribute_wrapper = soup.find("div", id="player_stats")

    attributes = attribute_wrapper.find_all("div", class_="column")

    for attribute in attributes:
        file.write(f"{attribute.find('h3').text} - ")
        table = attribute.find("table").find_all("tr")
        for row in table:
            name = row["id"]
            value = row.find_all("td")[1].text
            file.write(f"{name}: {value} -")

    file.write(",\n")


# Function to scrape the listing page and get player URLs
def scrape_listing_page(listing_url):
    response = requests.get(listing_url, timeout=10)
    if response.status_code != 200:
        return False

    soup = BeautifulSoup(response.text, "html.parser")

    print(soup)

    # Find all links to player pages - adjust selector as needed
    # Assuming player_links_wrapper is correctly finding the <div> containing player information
    player_links_wrapper = soup.find(
        "div", class_="players")  # This finds the <div>
    print(player_links_wrapper)

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
        scrape_player_data("https://fminside.net" + url, file)

    file.close()

    load_more_link = soup.find("a", class_="loadmore")


# Main URL to start scraping from - adjust as needed
listing_url = "https://fminside.net/beheer/modules/players/resources/inc/frontend/generate-player-table.php?ajax_request=1"
sequence_url = "https://fminside.net/beheer/modules/players/resources/inc/frontend/generate-player-table.php?ajax_request=1&loadmore=true"
scrape_listing_page(listing_url)
while True:
    exists = scrape_listing_page(sequence_url)
    if not exists:
        break

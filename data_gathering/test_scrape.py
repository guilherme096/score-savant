import requests
from bs4 import BeautifulSoup
import pyodbc
from datetime import datetime
import re

file = open("data.csv", "w")


def scrape_player_data(player_url, file):
    player_info = {}
    player_info["attributes"] = {}
    response = requests.get(player_url, timeout=10)
    soup = BeautifulSoup(response.text, "html.parser")

    # PLAYER NAME
    player_wrapper = soup.find("div", id="player")
    player_name_wrapper = soup.find("div", class_="title")  # Placeholder
    player_name = player_name_wrapper.find("h1").text

    player_info["name"] = player_name

    # PLAYER IMAGE
    player_image_style = player_wrapper.find("span").get("style")
    start = (
        player_image_style.find("url(") + 4
    )  # Find the start of the URL and adjust to skip 'url('
    end = player_image_style.find(
        ")", start
    )  # Find the end of the URL, starting from the end of 'url('
    url = player_image_style[start + 2 : end]
    player_info["url"] = url

    # file.write(f"Name: {player_name} - ")
    basic_info_wrapper = (
        player_name_wrapper.find("div", class_="meta").find("ul").find_all("li")
    )

    # CLUB AND COUNTRY
    club = basic_info_wrapper[0].find("a").find("span", class_="value").text
    country = basic_info_wrapper[1].find("span").find("a").text
    player_info["club"] = club
    player_info["nation"] = country

    # file.write(f"Club: {club} - Country: {country} - Image URL: {url} - ")

    player_info_wrapper = (
        player_wrapper.find("div", class_="column").find("ul").find_all("li")
    )

    for element in player_info_wrapper:
        field = element.find("span", class_="key").text
        if field == "Caps / Goals" or field == "Unique ID" or field == "Name":
            continue
        value = element.find("span", class_="value").text.replace(",", "/")
        if field == "Position(s)":
            value = value.split("/")[0]
            if value == "STST":
                value = "STC"
            elif value == "AMCAM":
                value = "CAM"
            elif value == "GKGK":
                value = "GK"
            elif value == "DCDC":
                value = "DC"
        player_info[field] = value
        # file.write(f"{field}: {value} - ")

    contract_wrapper = (
        player_wrapper.find_all("div", class_="column")[1].find("ul").find_all("li")
    )

    for element in contract_wrapper:
        field = element.find("span", class_="key").text
        value = element.find("span", class_="value").text
        if field == "Club":
            continue
        elif field == "Sell value":
            field = "Value"
            if value == "Not for sale":
                value = -1
        elif field == "Wages":
            value = float(value.split(" ")[1].replace(",", ""))
        elif field == "Contract end":
            # convert to DATE for sql
            field = "Contract end"
            value = datetime.strptime(value, "%Y-%m-%d").date()
        elif field == "Rel. clause":
            field = "release_clause"
            value = float(value.split(" ")[1].replace(",", ""))
        player_info[field] = value

    role_wrapper = (
        player_wrapper.find_all("div", class_="column")[2].find("ol").find_all("li")
    )

    role = role_wrapper[0].find("span", class_="key").text
    if "Attack" in role:
        role = role.split(" ")[:-1]
        role = re.sub(r"(?<!^)(?=[A-Z])", " ", role[0])
        role += " (At)"
        role = "".join(role)
    if "Support" in role:
        role = role.split(" ")[:-1]
        print(role)
        role = re.sub(r"(?<!^)(?=[A-Z])", " ", role[0])
        role += " (Su)"
        role = "".join(role)
    if "Defend" in role:
        role = role.split(" ")[:-1]
        name = re.sub(r"(?<!^)(?=[A-Z])", " ", role[0])
        role = " (De)"
        role = "".join(role)

    player_info["role"] = role

    attribute_wrapper = soup.find("div", id="player_stats")

    attributes = attribute_wrapper.find_all("div", class_="column")

    for attribute in attributes:
        # file.write(f"| {attribute.find('h3').text} - ")
        table = attribute.find("table").find_all("tr")
        for row in table:
            name = row["id"]
            value = row.find_all("td")[1].text
            # file.write(f"{name}: {value} - ")
            player_info["attributes"][name] = value

    # file.write(",\n")
    if "release_clause" not in player_info:
        player_info["release_clause"] = -1

    stored_procedure = (
        "EXEC dbo.AddPlayer ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?"
    )

    att_name_mapping = {
        "corners": "Corners",
        "crossing": "Crossing",
        "dribbling": "Dribbling",
        "finishing": "Finishing",
        "first-touch": "First_Touch",
        "free-kick-taking": "Free_Kick_Taking",
        "heading": "Heading",
        "long-shots": "Long_Shots",
        "long-throws": "Long_Throws",
        "marking": "Marking",
        "passing": "Passing",
        "penalty-taking": "Penalty_Taking",
        "tackling": "Tackling",
        "technique": "Technique",
        "aggression": "Aggression",
        "anticipation": "Anticipation",
        "bravery": "Bravery",
        "composure": "Composure",
        "concentration": "Concentration",
        "decisions": "Decisons",
        "determination": "Determination",
        "flair": "Flair",
        "leadership": "Leadership",
        "off-the-ball": "Off_The_Ball",
        "positioning": "Positioning",
        "teamwork": "Teamwork",
        "vision": "Vision",
        "work-rate": "Work_Rate",
        "acceleration": "Acceleration",
        "agility": "Agility",
        "balance": "Balance",
        "jumping-reach": "Jumping_Reach",
        "natural-fitness": "Natural_Fitness",
        "pace": "Pace",
        "stamina": "Stamina",
        "strength": "Strength",
    }

    print(player_info["attributes"].keys())
    attributes = [
        f"{att_name_mapping[key]}:{value}"
        for key, value in player_info["attributes"].items()
    ]
    attributes = ",".join(attributes)
    params = (
        player_info["name"],  # @name
        int(player_info["Age"]),  # @age
        int(player_info["Weight"][0]),  # @height
        int(player_info["Height"][0]),  # @weight
        player_info["nation"],  # @nation
        1,
        "Premier League",  # @league
        player_info["club"],  # @value
        player_info["Foot"],  # @position
        int(player_info["Value"]),  # @role
        player_info["Position(s)"],  # @contract_end
        player_info["role"],
        float(player_info["Wages"]),  # @release_clause
        player_info["Contract end"],  # @pace
        int(player_info["release_clause"]),  # @pace
        attributes,
        player_info["url"],
    )

    print(params)

    try:
        cursor = conn.cursor()
        cursor.execute(stored_procedure, params)
        conn.commit()
        print("Stored procedure executed successfully")
    except Exception as e:
        print(f"Error: {e}")
    finally:
        print("\n")
        cursor.close()


# Function to scrape the listing page and get player URLs
def scrape_listing_page(listing_url):
    response = requests.get(listing_url, timeout=10)
    if response.status_code != 200:
        return False

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
        scrape_player_data("https://fminside.net" + url, file)

    file.close()

    load_more_link = soup.find("a", class_="loadmore")


# Define your connection parameters
server = "mednat.ieeta.pt"
port = "8101"
database = "p5g5"  # Replace with your actual database name
username = "p5g5"
password = "bo_jack64"  # Replace with your actual password

# Create the connection string
connection_string = (
    f"DRIVER={{ODBC Driver 17 for SQL Server}};"
    f"SERVER={server},{port};"
    f"DATABASE={database};"
    f"UID={username};"
    f"PWD={password}"
)

# Establish the connection
try:
    conn = pyodbc.connect(connection_string)
    print("Connection successful")
except Exception as e:
    print(f"Error: {e}")


# Main URL to start scraping from - adjust as needed
listing_url = "https://fminside.net/beheer/modules/players/resources/inc/frontend/generate-player-table.php?ajax_request=1"
sequence_url = "https://fminside.net/beheer/modules/players/resources/inc/frontend/generate-player-table.php?ajax_request=1&loadmore=true"
url = "https://fminside.net/resources/inc/ajax/update_filter.php"
payload = {
    "page": "players",
    "database_version": "5",
    "name": "",
    "uid": "",
    "club": "",
    "nationality": "",
    "league": "Premier League",
    "min_age": "",
    "max_age": "",
    "max_value": "",
    "max_wage": "",
    "min_ability": "",
    "max_ability": "",
    "min_potential": "",
    "max_potential": "",
}

# Send the POST request
response = requests.post(url, data=payload)

# Print the response
print(f"Status Code: {response.status_code}")
print("Response Text:", response.text)
scrape_listing_page(listing_url)
while True:
    exists = scrape_listing_page(sequence_url)
    if not exists:
        break

conn.close()

USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [dbo].[Top5PlayersForEachLeague]
AS
WITH RankedPlayers AS
(
    SELECT
        p.player_id AS PlayerID,
        p.name AS PlayerName,
        l.league_id AS LeagueID,
        l.name AS LeagueName,
        cl.name AS ClubName,
        c.wage AS Wage,
        p.value AS Value,
        pos.name AS Position,
        ROW_NUMBER() OVER (PARTITION BY l.league_id ORDER BY p.player_id) AS RowNum
    FROM
        Player p
    INNER JOIN
        Club cl ON p.club_id = cl.club_id
    INNER JOIN
        League l ON cl.league_id = l.league_id
    INNER JOIN
        Contract c ON p.player_id = c.player_id
    INNER JOIN
        PlayerPosition pp ON p.player_id = pp.player_id
    INNER JOIN
        Position pos ON pp.position_id = pos.position_id
)
SELECT
    PlayerID,
    PlayerName,
    LeagueID,
    LeagueName,
    ClubName,
    Wage,
    Value,
    Position
FROM
    RankedPlayers
WHERE
    RowNum <= 5
GO

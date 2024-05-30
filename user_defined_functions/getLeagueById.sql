USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to compare a single league
CREATE FUNCTION dbo.getLeagueById
(
    @LeagueID INT
)
RETURNS TABLE
AS
RETURN
(
    SELECT
        l.league_id,
        l.name AS league_name,
        COUNT(DISTINCT c.club_id) AS total_clubs,
        COUNT(DISTINCT p.player_id) AS total_players,
        SUM(c.value_total) AS total_value,
        SUM(c.wage_total) AS total_wage,
        AVG(c.avg_att) AS avg_att
    FROM
        League l
        INNER JOIN Club c ON l.league_id = c.league_id
        INNER JOIN Player p ON c.club_id = p.club_id
    WHERE
        l.league_id = @LeagueID
    GROUP BY
        l.league_id,
        l.name
)
GO

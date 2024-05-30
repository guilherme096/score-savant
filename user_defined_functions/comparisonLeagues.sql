USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to compare two leagues
ALTER FUNCTION dbo.CompareLeagues
(
    @LeagueID1 INT,
    @LeagueID2 INT
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
        AVG(CASE WHEN oa.rating IS NOT NULL THEN oa.rating END) AS avg_outfield_rating,
        AVG(CASE WHEN ga.rating IS NOT NULL THEN ga.rating END) AS avg_goalkeeper_rating,
        AVG(CASE WHEN oa.rating IS NOT NULL THEN oa.rating ELSE ga.rating END) AS avg_combined_rating
    FROM
        League l
        INNER JOIN Club c ON l.league_id = c.league_id
        INNER JOIN Player p ON c.club_id = p.club_id
        LEFT JOIN PlayerRole pr ON p.player_id = pr.player_id
        LEFT JOIN RolePosition rp ON pr.role_position_id = rp.id
        LEFT JOIN Role r ON rp.role_position = r.role_id
        LEFT JOIN OutfieldAttributeRating oa ON p.player_id = oa.player_id
        LEFT JOIN GoalkeeperAttributeRating ga ON p.player_id = ga.player_id
    WHERE
        l.league_id IN (@LeagueID1, @LeagueID2)
    GROUP BY
        l.league_id,
        l.name
)
GO

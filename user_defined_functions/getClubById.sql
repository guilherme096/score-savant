USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get club information by club ID
CREATE FUNCTION dbo.GetClubByID
(
    @ClubID INT
)
RETURNS TABLE
AS
RETURN
(
    SELECT
        c.club_id,
        c.name AS club_name,
        l.name AS league_name,
        n.name AS nation_name,
        c.player_count,
        c.value_total,
        c.wage_total,
        c.value_average,
        c.wage_average,
        c.avg_att
    FROM
        Club c
    INNER JOIN
        League l ON c.league_id = l.league_id
    INNER JOIN
        Nation n ON c.nation_id = n.nation_id
    WHERE
        c.club_id = @ClubID
)
GO

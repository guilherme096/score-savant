USE [p5g5]

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

ALTER FUNCTION dbo.GetPlayerById
(
    @player_id INT
)
RETURNS TABLE
RETURN
    (
        SELECT
        p.*,
        c.wage,
        c.contract_end,
        c.duration,
        c.release_clause,

        cl.name AS club_name,

        n.name AS nation_name,

        l.name AS league_name

    FROM
        Player p
    INNER JOIN
        Contract c ON p.player_id = c.player_id
    INNER JOIN
        Club cl ON p.club_id = cl.club_id
    INNER JOIN
        Nation n ON p.nation_id = n.nation_id
    INNER JOIN
        League l ON cl.league_id = l.league_id
    WHERE
        p.player_id = @player_id
    )
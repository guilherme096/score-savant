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

        l.name AS league_name,

        oar_m.rating AS rating
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
    LEFT JOIN
        OutfieldAttributeRating oar_m ON p.player_id = oar_m.player_id
    LEFT JOIN
        Mental_Att a_m ON oar_m.att_id = a_m.att_id
    LEFT JOIN
        Physical_Att a_p ON oar_m.att_id= a_p.att_id
    LEFT JOIN
        Technical_Att a_t ON oar_m.att_id= a_t.att_id
    WHERE
        p.player_id = @player_id
    )

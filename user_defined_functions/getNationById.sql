USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get nation information by nation ID
CREATE FUNCTION dbo.GetNationByID
(
    @NationID INT
)
RETURNS @Result TABLE
(
    nation_id INT,
    nation_name NVARCHAR(255),
    total_leagues INT,
    league_names NVARCHAR(MAX),
    total_player_value DECIMAL(18, 2)
)
AS
BEGIN
    -- Insert aggregated results into the result table
    INSERT INTO @Result
    SELECT
        n.nation_id,
        n.name AS nation_name,
        COUNT(DISTINCT l.league_id) AS total_leagues,
        (
            SELECT STRING_AGG(name, ', ') AS league_names
            FROM (
                SELECT DISTINCT l.name
                FROM League l
                WHERE l.nation_id = n.nation_id
            ) AS sub
        ) AS league_names,
        SUM(p.value) AS total_player_value
    FROM
        Nation n
    LEFT JOIN
        League l ON n.nation_id = l.nation_id
    LEFT JOIN
        Player p ON p.nation_id = n.nation_id
    WHERE
        n.nation_id = @NationID
    GROUP BY
        n.nation_id, n.name;

    RETURN;
END;
GO

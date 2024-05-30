USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get league information with pagination, sorting, and search filters
CREATE FUNCTION dbo.GetLeaguesWithPagination
(
    @PageNumber INT,
    @PageSize INT,
    @OrderBy NVARCHAR(50) = NULL,
    @OrderDirection NVARCHAR(4) = NULL,
    @SearchLeagueName NVARCHAR(255) = NULL,
    @SearchNationName NVARCHAR(255) = NULL,
    @MinValueTotal DECIMAL(18,2) = NULL,
    @MaxValueTotal DECIMAL(18,2) = NULL
)
RETURNS TABLE
AS
RETURN
(
    SELECT 
        LeagueID,
        LeagueName,
        Nation,
        ValueTotal
    FROM
    (
        SELECT
            l.league_id AS LeagueID,
            l.name AS LeagueName,
            n.name AS Nation,
            SUM(c.value_total) AS ValueTotal,
            ROW_NUMBER() OVER (
                ORDER BY
                    CASE WHEN @OrderBy = 'LeagueName' AND @OrderDirection = 'ASC' THEN l.name END ASC,
                    CASE WHEN @OrderBy = 'LeagueName' AND @OrderDirection = 'DESC' THEN l.name END DESC,
                    CASE WHEN @OrderBy = 'Nation' AND @OrderDirection = 'ASC' THEN n.name END ASC,
                    CASE WHEN @OrderBy = 'Nation' AND @OrderDirection = 'DESC' THEN n.name END DESC,
                    CASE WHEN @OrderBy = 'ValueTotal' AND @OrderDirection = 'ASC' THEN SUM(c.value_total) END ASC,
                    CASE WHEN @OrderBy = 'ValueTotal' AND @OrderDirection = 'DESC' THEN SUM(c.value_total) END DESC,
                    -- Default order
                    l.league_id
            ) AS RowNum
        FROM
            League l
        INNER JOIN
            Nation n ON l.nation_id = n.nation_id
        LEFT JOIN
            Club c ON l.league_id = c.league_id
        WHERE
            (@SearchLeagueName IS NULL OR l.name LIKE '%' + @SearchLeagueName + '%')
            AND (@SearchNationName IS NULL OR n.name LIKE '%' + @SearchNationName + '%')
        GROUP BY
            l.league_id, l.name, n.name
        HAVING
            (@MinValueTotal IS NULL OR SUM(c.value_total) >= @MinValueTotal)
            AND (@MaxValueTotal IS NULL OR SUM(c.value_total) <= @MaxValueTotal)
    ) AS LeaguesWithRowNum
    WHERE
        RowNum > (@PageNumber - 1) * @PageSize
        AND RowNum <= @PageNumber * @PageSize
)
GO

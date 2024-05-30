USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

-- Create the UDF to get nation information with pagination, sorting, and search filters
CREATE FUNCTION dbo.GetNationsWithPagination
(
    @PageNumber INT,
    @PageSize INT,
    @OrderBy NVARCHAR(50) = NULL,
    @OrderDirection NVARCHAR(4) = NULL,
    @SearchNationName NVARCHAR(255) = NULL,
    @MinValueTotal DECIMAL(18,2) = NULL,
    @MaxValueTotal DECIMAL(18,2) = NULL
)
RETURNS TABLE
AS
RETURN
(
    SELECT 
        NationID,
        NationName,
        ValueTotal
    FROM
    (
        SELECT
            n.nation_id AS NationID,
            n.name AS NationName,
            SUM(p.value) AS ValueTotal,
            ROW_NUMBER() OVER (
                ORDER BY
                    CASE WHEN @OrderBy = 'NationName' AND @OrderDirection = 'ASC' THEN n.name END ASC,
                    CASE WHEN @OrderBy = 'NationName' AND @OrderDirection = 'DESC' THEN n.name END DESC,
                    CASE WHEN @OrderBy = 'ValueTotal' AND @OrderDirection = 'ASC' THEN SUM(p.value) END ASC,
                    CASE WHEN @OrderBy = 'ValueTotal' AND @OrderDirection = 'DESC' THEN SUM(p.value) END DESC,
                    -- Default order
                    n.nation_id
            ) AS RowNum
        FROM
            Nation n
        LEFT JOIN
            Player p ON p.nation_id = n.nation_id
        WHERE
            (@SearchNationName IS NULL OR n.name LIKE '%' + @SearchNationName + '%')
        GROUP BY
            n.nation_id, n.name
        HAVING
            (@MinValueTotal IS NULL OR SUM(p.value) >= @MinValueTotal)
            AND (@MaxValueTotal IS NULL OR SUM(p.value) <= @MaxValueTotal)
    ) AS NationsWithRowNum
    WHERE
        RowNum > (@PageNumber - 1) * @PageSize
        AND RowNum <= @PageNumber * @PageSize
)
GO

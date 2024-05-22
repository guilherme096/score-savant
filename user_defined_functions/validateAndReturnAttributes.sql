USE[p5g5]

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE FUNCTION dbo.VerifyAndReturnAttributes
(
    @Attributes NVARCHAR(MAX)
)
RETURNS @Result TABLE (Attribute NVARCHAR(255), Rating INT)
AS
BEGIN
    DECLARE @Pairs TABLE (Pair NVARCHAR(255))
    INSERT INTO @Pairs
    SELECT Item FROM dbo.SplitString(@Attributes, ',')

    DECLARE @Pair NVARCHAR(255), @Attribute NVARCHAR(255), @Rating INT

    DECLARE pair_cursor CURSOR FOR
    SELECT Pair FROM @Pairs

    OPEN pair_cursor
    FETCH NEXT FROM pair_cursor INTO @Pair

    WHILE @@FETCH_STATUS = 0
    BEGIN
        SET @Attribute = SUBSTRING(@Pair, 1, CHARINDEX(':', @Pair) - 1)
        SET @Rating = CAST(SUBSTRING(@Pair, CHARINDEX(':', @Pair) + 1, LEN(@Pair) - CHARINDEX(':', @Pair)) AS INT)

        IF EXISTS (SELECT 1 FROM Attribute WHERE name = @Attribute)
        BEGIN
            INSERT INTO @Result (Attribute, Rating)
            VALUES (@Attribute, @Rating)
        END

        FETCH NEXT FROM pair_cursor INTO @Pair
    END

    CLOSE pair_cursor
    DEALLOCATE pair_cursor

    RETURN
END

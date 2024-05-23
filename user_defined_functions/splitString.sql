USE[p5g5]

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE FUNCTION dbo.SplitString
(
    @Input NVARCHAR(MAX),
    @Delimiter CHAR(1)
)
RETURNS @Output TABLE (Item NVARCHAR(255))
AS
BEGIN
    DECLARE @Start INT, @End INT
    SELECT @Start = 1, @End = CHARINDEX(@Delimiter, @Input)
    WHILE @Start < LEN(@Input) + 1 BEGIN
        IF @End = 0  
            SET @End = LEN(@Input) + 1
        INSERT INTO @Output (Item)  
        VALUES(SUBSTRING(@Input, @Start, @End - @Start))
        SET @Start = @End + 1
        SET @End = CHARINDEX(@Delimiter, @Input, @Start)
    END
    RETURN
END

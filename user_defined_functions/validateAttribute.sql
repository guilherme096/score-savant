USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

ALTER FUNCTION [dbo].[ValidateAttributes]
(
    @attributes NVARCHAR(MAX)
)
RETURNS @ValidAttributes TABLE
(
    AttributeName NVARCHAR(255),
    AttributeRating INT,
    ErrorMessage NVARCHAR(255)
)
AS
BEGIN
    DECLARE @pos INT;
    DECLARE @str NVARCHAR(255);
    DECLARE @attribute_name NVARCHAR(255);
    DECLARE @attribute_rating INT;
    DECLARE @attribute_string NVARCHAR(MAX);

    SET @attribute_string = @attributes;

    -- Split attributes string and validate
    WHILE LEN(@attribute_string) > 0
    BEGIN
        SET @pos = CHARINDEX(',', @attribute_string);
        IF @pos = 0
            SET @pos = LEN(@attribute_string) + 1;
        SET @str = LEFT(@attribute_string, @pos - 1);
        SET @attribute_string = SUBSTRING(@attribute_string, @pos + 1, LEN(@attribute_string) - @pos);

        -- Split attribute name and rating
        SET @attribute_name = LEFT(@str, CHARINDEX(':', @str) - 1);
        SET @attribute_rating = CONVERT(INT, SUBSTRING(@str, CHARINDEX(':', @str) + 1, LEN(@str) - CHARINDEX(':', @str)));

        -- Validate attribute existence
        IF EXISTS (SELECT 1 FROM Attribute WHERE name = @attribute_name)
        BEGIN
            -- Insert valid attribute into the table variable
            INSERT INTO @ValidAttributes (AttributeName, AttributeRating, ErrorMessage)
            VALUES (@attribute_name, @attribute_rating, NULL);
        END
        ELSE
        BEGIN
            -- If attribute is not valid, return an error message
            INSERT INTO @ValidAttributes (AttributeName, AttributeRating, ErrorMessage)
            VALUES (NULL, NULL, 'Attribute not found: ' + @attribute_name);
            RETURN;
        END
    END

    RETURN;
END
GO

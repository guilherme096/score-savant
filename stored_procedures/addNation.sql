USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[AddNation]
    @nation NVARCHAR(255)
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @nation_id INT;

    -- Check if the nation already exists
    SELECT @nation_id = nation_id FROM Nation WHERE name = @nation;
    
    -- If the nation does not exist, insert it
    IF @nation_id IS NULL
    BEGIN
        INSERT INTO Nation (name) VALUES (@nation);
        SET @nation_id = SCOPE_IDENTITY();
    END
END
GO

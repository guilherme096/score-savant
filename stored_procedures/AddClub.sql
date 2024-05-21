USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[AddClub]
    @club NVARCHAR(255),
    @nation_id INT,
    @league_id INT
AS
BEGIN

   DECLARE @club_id INT;

   IF NOT EXISTS (SELECT 1 FROM League WHERE league_id = 0)
    BEGIN
        Raiserror('Nation id not valid');
    end

    SELECT @club_id = club_id FROM Club WHERE name = @club;
    IF @club_id IS NULL
    BEGIN
        INSERT INTO Club (name,league_id,nation_id) VALUES (@club, @league_id, @nation_id);
        SET @club_id = SCOPE_IDENTITY();
    END
END
GO

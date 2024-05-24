USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

ALTER PROCEDURE [dbo].[AddBasePlayer]
    @name NVARCHAR(255),
    @age INT,
    @weight INT,
    @height INT,
    @nation_id INT,
    @club_id INT,
    @foot NVARCHAR(255),
    @value INT,
    @player_type INT,
    @url NVARCHAR(MAX)
AS
BEGIN
    IF EXISTS(SELECT 1 FROM Player WHERE name = @name)
    BEGIN
        Raiserror('Player Already Exists',16,1);
    END

    IF NOT EXISTS(SELECT 1 FROM Club WHERE club_id = @club_id)
    BEGIN
        Raiserror('Club Doesnt Exist',16,1);
    END

    IF NOT EXISTS(SELECT 1 FROM Nation WHERE nation_id = @nation_id)
        BEGIN
            Raiserror('Nation Doesnt Exist',16,1);
        END

    IF @height < 0 OR @age < 0 OR @weight < 0 OR @value < 0 OR @foot LIKE ''
    BEGIN
        Raiserror('Invalid Values',16,1);
    END

    DECLARE @player_id INT;

    SELECT @player_id = player_id FROM Player WHERE name = @name
    IF @player_id IS NULL
    BEGIN
        INSERT INTO Player (name, age, weight, height, nation_id, club_id, foot, value, url) VALUES (@name, @age, @weight, @height, @nation_id, @club_id, @foot, @value, @url);
        SET @player_id = SCOPE_IDENTITY();

        IF @player_type = 0
        BEGIN
            INSERT INTO Outfield_Player (player_id) VALUES (@player_id);
        end
        ELSE
        BEGIN
            INSERT INTO Goalkeeper (player_id) VALUES (@player_id);
        end
    END
END
GO

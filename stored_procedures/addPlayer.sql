USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE dbo.AddPlayer
    @name NVARCHAR(255),
    @age INT,
    @weight INT,
    @height INT,
    @nation NVARCHAR(255),
    @nation_league_id INT,
    @league NVARCHAR(255),
    @club NVARCHAR(255),
    @foot NVARCHAR(255),
    @value INT,
    @position NVARCHAR(255),
    @role NVARCHAR(255),
    @wage DECIMAL(18,2),
    @contract_end DATE,
    @release_clause INT,
    @attributes NVARCHAR(MAX),
    @url NVARCHAR(MAX)
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @nation_id INT;
    DECLARE @league_id INT;
    DECLARE @club_id INT;
    DECLARE @player_id INT;
    DECLARE @position_id INT;
    DECLARE @role_id INT;
    DECLARE @player_type INT;

    -- Add or get Nation
    EXEC dbo.AddNation @nation = @nation;
    SELECT @nation_id = nation_id FROM Nation WHERE name = @nation;

    -- Add or get League
    EXEC dbo.AddLeague @league = @league, @nation = @nation_league_id;
    SELECT @league_id = league_id FROM League WHERE name = @league;

    -- Add or get Club
    EXEC dbo.AddClub @club = @club, @nation_id = @nation_league_id, @league_id = @league_id;
    SELECT @club_id = club_id FROM Club WHERE name = @club;

    -- Validate Position
    SET @position_id = dbo.ValidatePosition(@position);
    IF @position_id IS NULL
    BEGIN
        RAISERROR('Position not found: %s', 16, 1, @position);
        RETURN;
    END

    -- Determine player type based on position
    IF @position = 'GK'
    BEGIN
        SET @player_type = 1;
    END
    ELSE
    BEGIN
        SET @player_type = 0;
    END

    -- Validate Role
    SET @role_id = dbo.ValidateRole(@role);
    IF @role_id IS NULL
    BEGIN
        RAISERROR('Role not found: %s', 16, 1, @role);
        RETURN;
    END

    -- Add Base Player
    EXEC dbo.AddBasePlayer @name = @name, @age = @age, @weight = @weight, @height = @height, 
                           @nation_id = @nation_id, @club_id = @club_id, @foot = @foot, 
                           @value = @value, @player_type = @player_type, @url = @url;
    SELECT @player_id = player_id FROM Player WHERE name = @name;

    -- Add Contract
    EXEC dbo.AddContract @player_id = @player_id, @wage = @wage, @contract_end = @contract_end, @release_clause = @release_clause;

    -- Add Player Position
    EXEC dbo.AddPlayerPosition @position = @position_id, @player = @player_id;

    -- Add Player Role
    EXEC dbo.AddPlayerRole @role = @role_id, @player = @player_id;

    -- Add Player Attributes and Ratings
    IF @player_type = 0 -- Outfield Player
    BEGIN
        EXEC dbo.AddOutfieldAttributeRating @PlayerID = @player_id, @Attributes = @attributes;
    END
    ELSE -- Goalkeeper
    BEGIN
        EXEC dbo.AddGoalkeeperAttributeRating @PlayerID = @player_id, @Attributes = @attributes;
    END
END
GO

<<<<<<< HEAD
USE [p5g5]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE PROCEDURE [dbo].[AddPlayer]
    @name NVARCHAR(255),
    @weight INT,
    @height INT,
    @age INT,
    @club NVARCHAR(255),
    @nation NVARCHAR(255),
    @best_foot NVARCHAR(50),
    @position NVARCHAR(255),
    @value DECIMAL(18, 2),
    @player_role NVARCHAR(255),
    @wage DECIMAL(18, 2),
    @release_clause INT,
    @contract_end_date DATE,
    @duration INT,
    @attributes NVARCHAR(MAX) -- Expected format: "attribute1:rating1,attribute2:rating2,..."
AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @club_id INT;
    DECLARE @nation_id INT;
    DECLARE @position_id INT;
    DECLARE @role_id INT;
    DECLARE @player_id INT;
    DECLARE @attribute_name NVARCHAR(255);
    DECLARE @rating INT;
    DECLARE @pos INT;
    DECLARE @str NVARCHAR(255);
    DECLARE @player_type INT;
    DECLARE @attribute_rating INT;
    DECLARE @attribute_exists BIT;

    IF @position = 'GK'
    BEGIN
        SET @player_type = 0;
    END
    ELSE
    BEGIN
        SET @player_type = 1;
    END

    -- Get or insert club
    SELECT @club_id = club_id FROM Club WHERE name = @club;
    IF @club_id IS NULL
    BEGIN
        INSERT INTO Club (name) VALUES (@club);
        SET @club_id = SCOPE_IDENTITY();
    END

    -- Get or insert nation
    SELECT @nation_id = nation_id FROM Nation WHERE name = @nation;
    IF @nation_id IS NULL
    BEGIN
        INSERT INTO Nation (name) VALUES (@nation);
        SET @nation_id = SCOPE_IDENTITY();
    END

    -- Validate position
    SELECT @position_id = position_id FROM Position WHERE name = @position;
    IF @position_id IS NULL
    BEGIN
        RAISERROR('Position not found: %s', 16, 1, @position);
        RETURN;
    END

    -- Validate player role
    SELECT @role_id = role_id FROM Role WHERE name = @player_role;
    IF @role_id IS NULL
    BEGIN
        RAISERROR('Player role not found: %s', 16, 1, @player_role);
        RETURN;
    END

    -- Insert player
    INSERT INTO Player (name, age, weight, height, nation_id, club_id, foot, value) 
    VALUES (@name, @age, @weight, @height, @nation_id, @club_id, @best_foot, @value);
    SET @player_id = SCOPE_IDENTITY();

    -- Insert player role
    IF NOT EXISTS (SELECT 1 FROM PlayerRole WHERE player_id = @player_id AND role_position_id = @role_id)
    BEGIN
        INSERT INTO PlayerRole (player_id, role_position_id, rating) VALUES (@player_id, @role_id, @rating);
    END

    -- Insert player position
    IF NOT EXISTS (SELECT 1 FROM PlayerPosition WHERE player_id = @player_id AND position_id = @position_id)
    BEGIN
        INSERT INTO PlayerPosition (player_id, position_id) VALUES (@player_id, @position_id);
    END

    -- Insert contract
    INSERT INTO Contract (player_id, wage, duration, contract_end, release_clause) 
    VALUES (@player_id, @wage, @duration, @contract_end_date, @release_clause);

    -- Split attributes string and insert attributes
    WHILE LEN(@attributes) > 0
    BEGIN
        SET @pos = CHARINDEX(',', @attributes);
        IF @pos = 0
            SET @pos = LEN(@attributes) + 1;
        SET @str = LEFT(@attributes, @pos - 1);
        SET @attributes = SUBSTRING(@attributes, @pos + 1, LEN(@attributes) - @pos);

        -- Split attribute name and rating
        SET @attribute_name = LEFT(@str, CHARINDEX(':', @str) - 1);
        SET @attribute_rating = CONVERT(INT, SUBSTRING(@str, CHARINDEX(':', @str) + 1, LEN(@str) - CHARINDEX(':', @str)));

        -- Initialize the attribute exists flag
        SET @attribute_exists = 0;

        -- Validate attribute for outfield players
        IF @player_type = 1
        BEGIN
            IF EXISTS (SELECT 1 FROM Technical_Att WHERE att_id = @attribute_name)
            OR EXISTS (SELECT 1 FROM Mental_Att WHERE att_id = @attribute_name)
            OR EXISTS (SELECT 1 FROM Physical_Att WHERE att_id = @attribute_name)
            BEGIN
                SET @attribute_exists = 1;
                IF NOT EXISTS (SELECT 1 FROM OutfieldAttributeRating WHERE att_id = @attribute_name AND player_id = @player_id)
                BEGIN
                    INSERT INTO OutfieldAttributeRating (att_id, player_id, rating) VALUES (@attribute_name, @player_id, @attribute_rating);
                END
            END
        END
        -- Validate attribute for goalkeepers
        ELSE
        BEGIN
            IF EXISTS (SELECT 1 FROM Mental_Att WHERE att_id = @attribute_name)
            OR EXISTS (SELECT 1 FROM Physical_Att WHERE att_id = @attribute_name)
            OR EXISTS (SELECT 1 FROM Goalkeeping_Att WHERE att_id = @attribute_name)
            BEGIN
                SET @attribute_exists = 1;
                IF NOT EXISTS (SELECT 1 FROM GoalkeeperAttributeRating WHERE att_id = @attribute_name AND player_id = @player_id)
                BEGIN
                    INSERT INTO GoalkeeperAttributeRating (att_id, player_id, rating) VALUES (@attribute_name, @player_id, @attribute_rating);
                END
            END
        END

        -- If the attribute does not exist, raise an error
        IF @attribute_exists = 0
        BEGIN
            RAISERROR('Attribute not found: %s', 16, 1, @attribute_name);
            RETURN;
        END
    END

    SET NOCOUNT OFF;
END
GO

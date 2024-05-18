CREATE PROCEDURE AddPlayer
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
    @wage DECIMAL(18,2),
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
    DECLARE @contract_id INT;
    DECLARE @player_id INT;
    DECLARE @attribute_name NVARCHAR(255);
    DECLARE @attribute_id NVARCHAR(255);
    DECLARE @rating INT;
    DECLARE @pos INT;
    DECLARE @str NVARCHAR(255);
    DECLARE @player_type INT;
    DECLARE @attribute_rating INT;

    IF @position LIKE ('GK')
        BEGIN
            SET @player_type = 0;
        end
    ELSE
        SET @player_type = 1;


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

    -- Get or insert position
    SELECT @position_id = position_id FROM Position WHERE name = @position;
    IF @position_id IS NULL
    BEGIN
        INSERT INTO Position (name) VALUES (@position);
        SET @position_id = SCOPE_IDENTITY();
    END

    -- Get or insert player role


    -- Insert player
    INSERT INTO Player (name, age, weight, height, nation_id, club_id, foot, value) VALUES (@name, @age, @weight, @height, @nation_id, @club_id, @best_foot, @value);
    SET @player_id = SCOPE_IDENTITY();

    SELECT @role_id = role_id FROM Role WHERE name = @player_role;
    IF @role_id IS NULL
        BEGIN
            INSERT INTO Role (name) VALUES (@player_role);
            SET @role_id = SCOPE_IDENTITY();
        END

    -- Insert player role
    IF NOT EXISTS (SELECT 1 FROM PlayerRole WHERE player_id = @player_id AND role_position_id = @role_id)
        BEGIN
            INSERT INTO PlayerRole (player_id, role_position_id, rating) VALUES (@player_id, @role_id, @rating);
        END

    -- Insert contract
    INSERT INTO Contract (player_id, wage, duration, contract_end, release_clause) VALUES (@player_id, @wage, @duration, @contract_end_date, @release_clause);
    SET @contract_id = SCOPE_IDENTITY();

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

        -- Get or insert attribute
        SELECT @attribute_id = name FROM Attribute WHERE name = @attribute_name;
        IF @attribute_id IS NULL
        BEGIN
            INSERT INTO Attribute (name) VALUES (@attribute_name);
            SET @attribute_id = @attribute_name;
        END

        -- Insert player attribute
        IF @player_type = 0
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM GoalkeeperAttributeRating WHERE att_id = @attribute_id AND player_id = @player_id)
            BEGIN
                INSERT INTO GoalkeeperAttributeRating (att_id, player_id, rating) VALUES (@attribute_id, @player_id, @attribute_rating);
            END
        END
        ELSE
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM OutfieldAttributeRating WHERE att_id = @attribute_id AND player_id = @player_id)
            BEGIN
                INSERT INTO OutfieldAttributeRating (att_id, player_id, rating) VALUES (@attribute_id, @player_id, @attribute_rating);
            END
        END
    END

    SET NOCOUNT OFF;


END

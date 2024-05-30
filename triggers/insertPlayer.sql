USE [p5g5]
GO

CREATE TRIGGER trg_after_insert_players
ON Player
AFTER INSERT
AS
BEGIN
    DECLARE @club_id INT;
    DECLARE @total_players INT;
    DECLARE @total_wage DECIMAL(18, 2);
    DECLARE @total_value DECIMAL(18, 2);
    DECLARE @avg_att DECIMAL(18, 2);

    -- Get the club_id from the inserted row
    SELECT @club_id = club_id
    FROM INSERTED;

    -- Calculate the new totals
    SELECT @total_players = COUNT(*), 
           @total_wage = SUM(PC.wage), 
           @total_value = SUM(P.value)
    FROM Player P
    JOIN PlayerContract PC ON P.player_id = PC.player_id
    WHERE P.club_id = @club_id;

    -- Calculate the average attribute rating combining outfield and goalkeeper ratings
    SELECT @avg_att = AVG(CASE 
                            WHEN oa.rating IS NOT NULL THEN oa.rating 
                            ELSE ga.rating 
                          END)
    FROM Player P
    LEFT JOIN OutfieldAttributeRating oa ON P.player_id = oa.player_id
    LEFT JOIN GoalkeeperAttributeRating ga ON P.player_id = ga.player_id
    WHERE P.club_id = @club_id;

    -- Handle case where there are no players
    IF @total_players = 0 
    BEGIN
        SET @total_wage = 0;
        SET @total_value = 0;
        SET @avg_att = 0;
    END

    -- Update Club table
    UPDATE Club
    SET player_count = @total_players,
        wage_total = @total_wage,
        value_total = @total_value,
        wage_average = CASE WHEN @total_players > 0 THEN @total_wage / @total_players ELSE 0 END,
        value_average = CASE WHEN @total_players > 0 THEN @total_value / @total_players ELSE 0 END,
        avg_att = @avg_att
    WHERE club_id = @club_id;
END;
GO

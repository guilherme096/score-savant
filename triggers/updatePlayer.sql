USE [p5g5]
GO

CREATE TRIGGER trg_after_update_players
ON Player
AFTER UPDATE
AS
BEGIN
    DECLARE @old_club_id INT;
    DECLARE @new_club_id INT;
    DECLARE @total_players_old INT;
    DECLARE @total_wage_old DECIMAL(18, 2);
    DECLARE @total_value_old DECIMAL(18, 2);
    DECLARE @total_players_new INT;
    DECLARE @total_wage_new DECIMAL(18, 2);
    DECLARE @total_value_new DECIMAL(18, 2);

    -- Get the old and new club_id from the updated rows
    SELECT @old_club_id = DELETED.club_id, @new_club_id = INSERTED.club_id
    FROM DELETED
    JOIN INSERTED ON DELETED.player_id = INSERTED.player_id;

    -- If the club has changed, update the old club's totals first
    IF @old_club_id IS NOT NULL AND @old_club_id <> @new_club_id
    BEGIN
        -- Calculate the new totals for the old club
        SELECT @total_players_old = COUNT(*), 
               @total_wage_old = SUM(PC.wage), 
               @total_value_old = SUM(PC.value)
        FROM Player P
        JOIN PlayerContract PC ON P.player_id = PC.player_id
        WHERE P.club_id = @old_club_id;

        -- Handle case where there are no players in the old club
        IF @total_players_old = 0 
        BEGIN
            SET @total_wage_old = 0;
            SET @total_value_old = 0;
        END

        -- Update the old club's totals
        UPDATE Club
        SET player_count = @total_players_old,
            wage_total = @total_wage_old,
            value_total = @total_value_old,
            wage_average = CASE WHEN @total_players_old > 0 THEN @total_wage_old / @total_players_old ELSE 0 END,
            value_average = CASE WHEN @total_players_old > 0 THEN @total_value_old / @total_players_old ELSE 0 END
        WHERE club_id = @old_club_id;
    END

    -- Calculate the new totals for the new club
    SELECT @total_players_new = COUNT(*), 
           @total_wage_new = SUM(PC.wage), 
           @total_value_new = SUM(PC.value)
    FROM Player P
    JOIN PlayerContract PC ON P.player_id = PC.player_id
    WHERE P.club_id = @new_club_id;

    -- Handle case where there are no players in the new club
    IF @total_players_new = 0 
    BEGIN
        SET @total_wage_new = 0;
        SET @total_value_new = 0;
    END

    -- Update the new club's totals
    UPDATE Club
    SET player_count = @total_players_new,
        wage_total = @total_wage_new,
        value_total = @total_value_new,
        wage_average = CASE WHEN @total_players_new > 0 THEN @total_wage_new / @total_players_new ELSE 0 END,
        value_average = CASE WHEN @total_players_new > 0 THEN @total_value_new / @total_players_new ELSE 0 END
    WHERE club_id = @new_club_id;
END;

CREATE PROCEDURE [dbo].[AddContract]
    @player_id INT,
    @wage decimal(18,2),
    @contract_end date,
    @release_clause INT
AS
BEGIN
    DECLARE @duration INT;
    DECLARE @currentYear INT
    DECLARE @endYear INT

    -- Get the current year
    SET @currentYear = YEAR(GETDATE())

    -- Get the year from the input date
    SET @endYear = YEAR(@contract_end)

    -- Calculate the difference
    SET @duration = @endYear - @currentYear

    IF EXISTS(SELECT 1 FROM Contract WHERE player_id = @player_id)
        BEGIN
            Raiserror('Player Already Has Contract',16,1);
        END

    IF @wage < 0 OR @release_clause < -2 OR @contract_end IS NULL
        BEGIN
            Raiserror('Invalid Values',16,1);
        END


    INSERT INTO Contract (player_id, wage, duration, contract_end, release_clause) VALUES (@player_id, @wage, @duration, @contract_end, @release_clause);
END
go


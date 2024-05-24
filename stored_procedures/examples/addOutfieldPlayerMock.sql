USE [p5g5]
GO

-- Example call to AddPlayer stored procedure
EXEC dbo.AddPlayer
    @name = 'Nuno Santos',
    @age = 25,
    @weight = 75,
    @height = 180,
    @nation = 'England',
    @nation_league_id = 1,
    @league = 'Premier League',
    @club = 'LA Galaxy',
    @foot = 'Right',
    @value = 1000000,
    @player_type = 0, -- 0 for outfield player, 1 for goalkeeper
    @position = 'STC',
    @role = 'Complete Forward (At)',
    @wage = 50000,
    @contract_end = '2026-06-30',
    @release_clause = 5000000,
    @attributes = 'Corners:10,Crossing:13,Dribbling:16,Finishing:18,First_Touch:17,Free_Kick_Taking:9,Heading:16,Long_Shots:15,Long_Throws:2,Marking:6,Passing:14,Penalty_Taking:18,Tackling:3,Technique:16,Aggression:10,Anticipation:18,Bravery:17,Composure:18,Concentration:15,Decisons:16,Determination:20,Flair:13,Leadership:15,Off_The_Ball:19,Positioning:3,Teamwork:17,Vision:15,Work_Rate:20,Acceleration:17,Agility:14,Balance:14,Jumping_Reach:16,Natural_Fitness:20,Pace:18,Stamina:20,Strength:17',
    @url = ''

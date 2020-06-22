DELIMITER $$
CREATE OR REPLACE PROCEDURE create_game
(
    IN name varchar(255), 
    IN version varchar(20),
    IN number_of_players int,
    IN investor bit,
    IN tax bit,
    IN startingMode varchar(20)
)
BEGIN
    DECLARE game_id bigint unsigned;
    SELECT uuid_short() INTO game_id;

    
	INSERT INTO game 
	(game_id, name, version, number_of_players, investor_card, tax_increase_only_bonus, starting_mode)
	VALUES
	(game_id, name, version, number_of_players, investor, tax, startingMode);

    SELECT game_id;

END $$
DELIMITER ;
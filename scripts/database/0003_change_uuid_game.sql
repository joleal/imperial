
DELIMITER $
BEGIN NOT ATOMIC
    IF EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS 
        WHERE COLUMN_NAME = 'Game_id' 
        AND TABLE_SCHEMA = 'imperial'
        AND DATA_TYPE = 'bigint') 
    THEN
        ALTER TABLE game CHANGE COLUMN game_id game_id_old bigint NOT NULL;
        ALTER TABLE game ADD game_id VARCHAR(64) AFTER game_id_old;
        UPDATE game SET game_id = CAST(game_id_old AS VARCHAR(64));
        ALTER TABLE game DROP game_id_old;
        ALTER TABLE game CHANGE COLUMN game_id game_id varchar(64) not null primary key;

        ALTER TABLE game_player CHANGE COLUMN fk_game fk_game_old bigint NOT NULL;
        ALTER TABLE game_player ADD fk_game varchar(64) AFTER fk_game_old;
        UPDATE game_player SET fk_game = CAST(fk_game_old AS VARCHAR(64));
        ALTER TABLE game_player DROP fk_game_old;
        ALTER TABLE game_player CHANGE COLUMN fk_game fk_game varchar(64) not null;

    END IF;
END $
DELIMITER ;
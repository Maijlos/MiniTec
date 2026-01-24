SET FOREIGN_KEY_CHECKS = 0;

-- 1. Revert Station
ALTER TABLE Station DROP FOREIGN KEY fk_station_project;
ALTER TABLE Station ADD CONSTRAINT Station_ibfk_1 
    FOREIGN KEY (project_id) REFERENCES Project(id);

-- 2. Revert State
ALTER TABLE State DROP FOREIGN KEY fk_state_station;
ALTER TABLE State ADD CONSTRAINT State_ibfk_1 
    FOREIGN KEY (station_id) REFERENCES Station(id);

-- 3. Revert User_Project
ALTER TABLE User_Project DROP FOREIGN KEY fk_user_project_project;
ALTER TABLE User_Project ADD CONSTRAINT User_Project_ibfk_2 
    FOREIGN KEY (project_id) REFERENCES Project(id);

SET FOREIGN_KEY_CHECKS = 1;
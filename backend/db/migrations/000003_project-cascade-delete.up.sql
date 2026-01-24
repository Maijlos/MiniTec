-- Disable foreign key checks to prevent errors during the swap
SET FOREIGN_KEY_CHECKS = 0;

-- 1. Update Station
ALTER TABLE Station DROP FOREIGN KEY Station_ibfk_1;
ALTER TABLE Station ADD CONSTRAINT fk_station_project 
    FOREIGN KEY (project_id) REFERENCES Project(id) 
    ON DELETE CASCADE;

-- 2. Update State
ALTER TABLE State DROP FOREIGN KEY State_ibfk_1;
ALTER TABLE State ADD CONSTRAINT fk_state_station 
    FOREIGN KEY (station_id) REFERENCES Station(id) 
    ON DELETE CASCADE;

-- 3. Update User_Project
ALTER TABLE User_Project DROP FOREIGN KEY User_Project_ibfk_2;
ALTER TABLE User_Project ADD CONSTRAINT fk_user_project_project 
    FOREIGN KEY (project_id) REFERENCES Project(id) 
    ON DELETE CASCADE;

SET FOREIGN_KEY_CHECKS = 1;
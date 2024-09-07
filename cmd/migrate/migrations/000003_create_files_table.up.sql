CREATE TABLE IF NOT EXISTS files (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL, 
  `folder_id` INT UNSIGNED DEFAULT NULL, 
  
  `file_name` VARCHAR(255) NOT NULL,
  `file_size` BIGINT UNSIGNED DEFAULT 0, 
  `file_type` VARCHAR(50) DEFAULT 'unknown',
  `file_path` VARCHAR(255) NOT NULL, 
  `description` TEXT,
  
  `mime_type` VARCHAR(100) DEFAULT 'application/octet-stream', 
  
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  
  PRIMARY KEY (`id`),
  
  CONSTRAINT `fk_files_user`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE,
  
  CONSTRAINT `fk_files_folder`
    FOREIGN KEY (`folder_id`) REFERENCES `folders`(`id`)
    ON DELETE SET NULL 
);

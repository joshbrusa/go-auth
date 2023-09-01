CREATE TABLE post (
  ID INT AUTO_INCREMENT PRIMARY KEY,
  CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  Title VARCHAR(255) NOT NULL,
  Body TEXT NOT NULL
);
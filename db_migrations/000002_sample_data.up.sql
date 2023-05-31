INSERT INTO roles (id, name, created_at) VALUES
(1, 'ROLE_ADMIN', CURRENT_TIMESTAMP),
(2, 'ROLE_MODERATOR', CURRENT_TIMESTAMP),
(3, 'ROLE_USER', CURRENT_TIMESTAMP)
;

INSERT INTO users (email, password, name, created_at) VALUES
('admin@gmail.com', 'admin', 'Admin', CURRENT_TIMESTAMP),
('demo@gmail.com', 'demo', 'Demo User', CURRENT_TIMESTAMP),
('siva@gmail.com', 'siva', 'Siva', CURRENT_TIMESTAMP)
;

INSERT INTO user_role (user_id, role_id) VALUES
(1, 1),
(1, 2),
(1, 3),
(2, 3),
(3, 2)
;

insert into posts(url, title, content, created_by, created_at) values
('https://linuxize.com/post/how-to-remove-docker-images-containers-volumes-and-networks/','How To Remove Docker Containers, Images, Volumes, and Networks','How To Remove Docker Containers, Images, Volumes, and Networks',1,CURRENT_TIMESTAMP),
('https://reflectoring.io/unit-testing-spring-boot/','All You Need To Know About Unit Testing with Spring Boot','All You Need To Know About Unit Testing with Spring Boot',1,CURRENT_TIMESTAMP),
('https://blog.jooq.org/2014/06/25/flyway-and-jooq-for-unbeatable-sql-development-productivity/','Flyway and jOOQ for Unbeatable SQL Development Productivity','Flyway and jOOQ for Unbeatable SQL Development Productivity',1,CURRENT_TIMESTAMP),
('https://www.marcobehler.com/guides/java-microservices-a-practical-guide','Java Microservices: A Practical Guide','Java Microservices: A Practical Guide',1,CURRENT_TIMESTAMP),
('https://sivalabs.in/2020/02/spring-boot-integration-testing-using-testcontainers-starter/','SpringBoot Integration Testing using TestContainers Starter','SpringBoot Integration Testing using TestContainers Starter',1,CURRENT_TIMESTAMP),
('https://medium.com/faun/continuous-integration-of-java-project-with-github-actions-7a8a0e8246ef','Continuous Integration of Java project with GitHub Actions','Continuous Integration of Java project with GitHub Actions',1,CURRENT_TIMESTAMP)
;
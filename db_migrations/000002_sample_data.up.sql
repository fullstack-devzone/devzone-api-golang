INSERT INTO users (id, name, email, password, role, created_at) VALUES
(1, 'Admin','admin@gmail.com','$2a$14$rPRrJzfPp0P4jMVF7EI.yeWWL8hfVifKkt2BJQChgC7qDemh8A5Q2', 'ROLE_ADMIN', CURRENT_TIMESTAMP),
(2, 'Siva','siva@gmail.com', '$2a$14$Nx4oY3HsxMCl9O.nb4FvseO8xaf1Lx/f7WsJf2ziIYAVDYWbA/Cru', 'ROLE_USER', CURRENT_TIMESTAMP),
(3, 'Demo','demo@gmail.com', '$2a$14$HRg1m3oqH64qM7gfqBHlOumqDtOT9DED9TvRK2NDXfEqT3Grx8bv6', 'ROLE_USER', CURRENT_TIMESTAMP)
;

insert into posts(id, url, title, content, created_by, created_at) values
(1, 'https://linuxize.com/post/how-to-remove-docker-images-containers-volumes-and-networks/','How To Remove Docker Containers, Images, Volumes, and Networks','How To Remove Docker Containers, Images, Volumes, and Networks',1,CURRENT_TIMESTAMP),
(2, 'https://reflectoring.io/unit-testing-spring-boot/','All You Need To Know About Unit Testing with Spring Boot','All You Need To Know About Unit Testing with Spring Boot',1,CURRENT_TIMESTAMP),
(3, 'https://blog.jooq.org/2014/06/25/flyway-and-jooq-for-unbeatable-sql-development-productivity/','Flyway and jOOQ for Unbeatable SQL Development Productivity','Flyway and jOOQ for Unbeatable SQL Development Productivity',1,CURRENT_TIMESTAMP),
(4, 'https://www.marcobehler.com/guides/java-microservices-a-practical-guide','Java Microservices: A Practical Guide','Java Microservices: A Practical Guide',1,CURRENT_TIMESTAMP),
(5, 'https://sivalabs.in/2020/02/spring-boot-integration-testing-using-testcontainers-starter/','SpringBoot Integration Testing using TestContainers Starter','SpringBoot Integration Testing using TestContainers Starter',1,CURRENT_TIMESTAMP),
(6, 'https://medium.com/faun/continuous-integration-of-java-project-with-github-actions-7a8a0e8246ef','Continuous Integration of Java project with GitHub Actions','Continuous Integration of Java project with GitHub Actions',1,CURRENT_TIMESTAMP),
(7, 'https://www.sivalabs.in/spring-boot-3-error-reporting-using-problem-details','Spring Boot 3 : Error Responses using Problem Details for HTTP APIs','Spring Boot 3 : Error Responses using Problem Details for HTTP APIs',1,CURRENT_TIMESTAMP),
(8, 'https://www.sivalabs.in/using-java-records-with-spring-boot-3','Using Java Records with Spring Boot 3','Using Java Records with Spring Boot 3',1,CURRENT_TIMESTAMP),
(9, 'https://www.sivalabs.in/how-to-not-to-ask-for-technical-help','How (not) to ask for Technical Help?','How (not) to ask for Technical Help?',1,CURRENT_TIMESTAMP),
(10, 'https://www.sivalabs.in/how-springboot-autoconfiguration-magic','How SpringBoot AutoConfiguration magic works?','How SpringBoot AutoConfiguration magic works?',1,CURRENT_TIMESTAMP),
(11, 'https://www.sivalabs.in/why-springboot','Why SpringBoot?','Why SpringBoot?',1,CURRENT_TIMESTAMP),
(12, 'https://www.sivalabs.in/2013/12/clean-code-dont-mix-different-levels-of-abstractions','Clean Code: Don’t mix different levels of abstractions','Clean Code: Don’t mix different levels of abstractions',1,CURRENT_TIMESTAMP)
;
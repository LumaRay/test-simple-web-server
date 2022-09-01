package com.example.testspringweb;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;

import java.util.Arrays;

@SpringBootApplication
public class HelloSpringWebApplication {

	public static void main(String[] args) {
		SpringApplication.run(HelloSpringWebApplication.class, args);
	}

	@Bean
	public CommandLineRunner commandLineRunner(ApplicationContext ctx) {
		return args -> {

//			System.out.println("Let's inspect the beans provided by Spring Boot:");
//
//			String[] beanNames = ctx.getBeanDefinitionNames();
//			Arrays.sort(beanNames);
//			for (String beanName : beanNames) {
//				System.out.println(beanName);
//			}

		};
	}
}


//import org.springframework.boot.SpringApplication;
//import org.springframework.boot.autoconfigure.SpringBootApplication;
//import org.springframework.web.bind.annotation.GetMapping;
//import org.springframework.web.bind.annotation.RestController;
//
//// This annotation instructs Spring to initialize its configuration - which is needed to start a
//// new application
//@SpringBootApplication
//// Indicates that this class contains RESTful methods to handle incoming HTTP requests
//@RestController
//public class HelloSpringWebApplication {
//
//	// We can start our application by calling the run method with the primary class
//	public static void main(String[] args) {
//		SpringApplication.run(HelloSpringWebApplication.class, args);
//	}
//
//	// The `GetMapping` annotation indicates that this method should be called
//	// when handling GET requests to the "/simple-request" endpoint
//	@GetMapping("/")
//	public String simpleRequest() {
//		// In this case, we return the plain text response "ok"
//		return "ok";
//	}
//}
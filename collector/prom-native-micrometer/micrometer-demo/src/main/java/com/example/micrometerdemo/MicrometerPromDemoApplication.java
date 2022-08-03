package com.example.micrometerdemo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.scheduling.annotation.EnableScheduling;

@SpringBootApplication
@EnableScheduling
public class MicrometerPromDemoApplication {

    public static void main(String[] args) {
        SpringApplication.run(MicrometerPromDemoApplication.class, args);
    }

}

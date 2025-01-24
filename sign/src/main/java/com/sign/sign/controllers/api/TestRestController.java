package com.sign.sign.controllers.api;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;


@RestController
public class TestRestController {
    
    @GetMapping("/test")
    @ResponseBody
    public String test() {
        RestTemplate restTemplate = new RestTemplate();
        String goApiUrl = "http://localhost:8081/api/get/users";
        String response = restTemplate.getForObject(goApiUrl, String.class);
        return response;
    }

    @GetMapping("/test2")
    public ResponseEntity<String> test2() {
        RestTemplate restTemplate = new RestTemplate();
        String goApiUrl = "http://localhost:8081/api/get/users";
        String response = restTemplate.getForObject(goApiUrl, String.class);

        HttpHeaders headers = new HttpHeaders();
        headers.add(HttpHeaders.CONTENT_TYPE, "application/json");

        return new ResponseEntity<>(response, headers, HttpStatus.OK);
    }
}

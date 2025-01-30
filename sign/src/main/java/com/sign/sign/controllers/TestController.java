package com.sign.sign.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;



@Controller
public class TestController {
    @GetMapping("/index")
    public String index() {
        return "index";
    }

    @GetMapping("/sign")
    public String sign() {
        return "sign";
    }

    @GetMapping("/login")
    public String login() {
        return "login";
    }

    @GetMapping("/signup")
    public String signup() {
        return "signup";
    }
    
    @GetMapping("/verification")
    public String verification() {
        return "verification";
    }
    
    @GetMapping("/keyform")
    public String keyform() {
        return "keyform";
    }   
}

package com.sign.sign;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import io.swagger.v3.oas.models.Components;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Info;

@Configuration
public class SwaggerConfig {
    
    @Bean
    public OpenAPI openAPI() {
        return new OpenAPI()
                .components(new Components())   // Swagger 구성 요소 설정
                .info(apiInfo());               // API 정보 설정
    }

    private Info apiInfo() {
        return new Info().title("Sign Swagger")          // API 문서 제목
                        .description("Sign REST API")    // API 문서 설명
                        .version("1.0.0");                  // API 문서 버전
    }
}

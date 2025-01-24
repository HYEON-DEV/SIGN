package com.sign.sign;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import com.sign.sign.interceptors.MyInterceptor;

@Configuration
@SuppressWarnings("null")
public class MyWebConfig implements WebMvcConfigurer {

    // MyInterceptor 객체를 자동 주입 받는다
    // 이 과정에서 myInterceptor @Autowired로 선언된 UtilHelper 객체도 자동 주입된다
    @Autowired
    private MyInterceptor myInterceptor;

    
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // 직접 정의한 MyInterceptor를 Spring에 등록
        InterceptorRegistration ir = registry.addInterceptor(myInterceptor);

        // 해당 경로는 인터셉터가 가로채지 않는다
        ir.excludePathPatterns("/error", "/robots.txt", "/favicon.ico", "assets/**");
    }

    /** 설정파일에 명시된 업로드 저장경로와 URL상의 경로를 맵핑 시킨다 */
    @Override
    public void addResourceHandlers(ResourceHandlerRegistry registry) {
        // registry.addResourceHandler(String.format("%s/**", uploadUrl)).addResourceLocations(String.format("file://%s/", uploadDir));

        registry.addResourceHandler("/static/**").addResourceLocations("classpath:/static/");
    }
    
}

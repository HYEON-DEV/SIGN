package com.sign.sign.controllers.api;

import java.io.IOException;

import org.springframework.core.io.FileSystemResource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.http.client.MultipartBodyBuilder;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.reactive.function.client.WebClient;
import org.springframework.web.reactive.function.client.WebClientResponseException;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Mono;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;


@Slf4j
@RestController
public class TestRestController2 {
    
    /**
     * WebClient.Builder를 사용하여 WebClient 인스턴스 생성
     * WebClient는 HTTP 요청을 보내고 응답을 받는 데 사용되는 클라이언트
     * 기본 URL을 http://localhost:8081로 설정한다.
     */

    private final WebClient webClient;

    // 생성자에서 WebClient.Builder를 사용하여 WebClient 인스턴스 생성
    public TestRestController2(WebClient.Builder webClientBuilder) {
        // build 메서드를 호출하여 WebClient 인스턴스를 생성
        this.webClient = webClientBuilder.baseUrl("http://localhost:8081").build();
        // 이 기본 URL은 Go 서버의 주소
        // 클래스의 멤버 변수인 webClient에 생성된 WebClient 인스턴스 할당
    }

    @GetMapping("/test3")
    public Mono<ResponseEntity<String>> test3() {
        return this.webClient.get()     // WebClient 사용한 GET 요청
                    .uri("/api/get/users")      // 요청할 엔드포인트 설정
                    .retrieve()     // 요청을 실행하고 응답 받는다
                    .bodyToMono(String.class)   // Mono<String>으로 변환 후
                    .map( response -> {
                        HttpHeaders headers = new HttpHeaders();
                        headers.add(HttpHeaders.CONTENT_TYPE, "application/json");
                        return new ResponseEntity<>(response, headers, HttpStatus.OK);
                        // ResponseEntity로 랩핑하여 반환
                    } );
    }


    /**
     * ECDSA 키 생성
     * @return JSON
     */
    @PostMapping("/api/generate_key")
    public Mono<ResponseEntity<String>> generateKey() {
        return this.webClient.post()
                    .uri("/api/generate_key")
                    .retrieve()
                    .bodyToMono(String.class)
                    .map( response -> {
                        HttpHeaders headers = new HttpHeaders();
                        headers.add(HttpHeaders.CONTENT_TYPE, "application/json");
                        return new ResponseEntity<>(response, headers, HttpStatus.OK);
                    });
    }


    /**
     * 서명 생성
     * @return
     */
    @PostMapping("/api/generate_sign")
    public Mono<String> generateSign(@RequestParam("input_privatekey") MultipartFile privateKeyFile) {
        MultipartBodyBuilder builder = new MultipartBodyBuilder();
        builder.part("input_privatekey", privateKeyFile.getResource());

        return webClient.post()
                .uri("/api/generate_sign")
                .header(HttpHeaders.CONTENT_TYPE, MediaType.MULTIPART_FORM_DATA_VALUE)
                .bodyValue(builder.build())
                .retrieve()
                .bodyToMono(String.class);
    }


    // @PostMapping("/api/generate_sign")
    // public Mono<ResponseEntity<String>> generateSign(@RequestParam("privateKeyFile") MultipartFile privateKeyFile) {
    //     // return this.webClient.post()
    //     //             .uri("/api/generate_sign")
    //     //             .contentType(MediaType.MULTIPART_FORM_DATA)
    //     //             .bodyValue(createMultipartBody(privateKeyFile,"privateKeyFile")) // 파일을 multipart로 전달
    //     //             .retrieve()
    //     //             .bodyToMono(String.class)
    //     //             .map( response -> {
    //     //                 HttpHeaders headers = new HttpHeaders();
    //     //                 headers.add(HttpHeaders.CONTENT_TYPE, "application/json");
    //     //                 return new ResponseEntity<>(response, headers, HttpStatus.OK);
    //     //             });

    //     MultiValueMap<String, Object> body = createMultipartBody(privateKeyFile, "privateKeyFile");

    //     return this.webClient.post()
    //             .uri("/api/generate_sign")
    //             .bodyValue(body)
    //             .retrieve()
    //             .bodyToMono(String.class)
    //             .map(response -> {
    //                 HttpHeaders headers = new HttpHeaders();
    //                 headers.add(HttpHeaders.CONTENT_TYPE, "application/json");
    //                 return new ResponseEntity<>(response, headers, HttpStatus.OK);
    //             });
    // }

    // // Java 클라이언트에서 파일을 Go 서버에 전송하는 방법은 WebClient를 사용하여 multipart/form-data 형식으로 파일을 전송하는 것
    // // MultipartFile을 Go 서버에 전송할 수 있는 multipart/form-data 형식으로 변환하는 메서드
    // private MultiValueMap<String, Object> createMultipartBody(MultipartFile file, String inputFile) {
    //     MultiValueMap<String, Object> body = new LinkedMultiValueMap<>();
    //     // MultipartFile을 실제 파일로 변환
    //     // WebClient의 bodyValue에 MultiValueMap을 전달하여 파일 전송
    //     try {
    //         body.add(inputFile, new FileSystemResource(file.getResource().getFile()));
    //     } catch (IOException e) {
    //         log.error("MultipartFile 변환 실패", e);
    //     }   
    //     return body;
    // }

    
    /**
     * Spring Boot에서는 파일을 받아서 WebClient를 통해 Go 서버로 전송합니다. 이때 MultipartBodyBuilder를 사용하여 파일을 multipart/form-data 형식으로 전달합니다.
Go 서버에서는 파일을 읽고, JSON으로 파싱한 후, 서명을 생성하는 로직을 처리합니다.
뷰에서는 사용자가 업로드한 파일을 FormData로 감싸서 POST 요청을 보내고, 서버에서 처리한 후 결과를 반환합니다.
    */
}

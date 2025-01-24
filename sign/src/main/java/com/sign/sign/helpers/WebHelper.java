package com.sign.sign.helpers;

import java.net.URLEncoder;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.io.PrintWriter;
import java.io.UnsupportedEncodingException;

import jakarta.servlet.http.Cookie;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Component  //  <- 스프링에게 이 클래스가 빈(Bean)임을 알려준다
public class WebHelper {

    @Autowired
    private HttpServletRequest request;

    @Autowired
    private HttpServletResponse response;
    
    /**
     * 클라이언트의 IP 주소 가져오는 메서드
     * @return IP 주소
     */
    public String getClientIp() {
        String ip = request.getHeader("X-Forwarded-For");

        if (ip == null) {
            ip = request.getHeader("Proxy-Client-IP");
        }
        if (ip == null) {
            ip = request.getHeader("WL-Proxy-Client-IP"); // 웹로직
        }
        if (ip == null) {
            ip = request.getHeader("HTTP_CLIENT_IP");
        }
        if (ip == null) {
            ip = request.getHeader("HTTP_X_FORWARDED_FOR");
        }
        if (ip == null) {
            ip = request.getRemoteAddr();
        }

        return ip;
    }

    
    /**
     * 쿠키값을 저장한다
     * @param name - 쿠키 이름
     * @param value - 쿠키 값
     * @param maxAge - 쿠키 유효 시간 (0이면 지정 안함, 음수일 경우 즉시 삭제)
     * @param domain - 쿠키 도메인
     * @param path - 쿠키 경로
     */
    public void writeCookie(String name, String value, int maxAge, String domain, String path) throws Exception {
        if ( value!=null && !value.equals("") ) {
            try {
                // import java.net.URLEncoder;
                value = URLEncoder.encode(value, "utf-8");
            } catch (UnsupportedEncodingException e) {
                // import java.io.UnsupportedEncodingException;
                log.error("쿠키값 인코딩 실패", e);
                throw e;
            }
        }

        /** 2) 쿠키 저장하기
         * 저장할 쿠키 객체 생성
         * import jakarta.servlet.http.Cookie;
         */
        Cookie cookie = new Cookie(name, value);

        // 유효경로 => 사이트 전역에 대한 설정
        //cookie.setPath("/");
        cookie.setPath(path);
        
        /**
         * 유효 도메인 (로컬 개발시에는 설정할 필요 X)
         * => "www.naver.com" 인 경우, ".naver.com"으로 설정 
         * cookie.setDomain("localhost");
         */
        if (domain != null) {
            cookie.setDomain(domain);
        }

        // 유효시간 설정 ( 0이하면 즉시삭제, 초단위 )
        // 설정하지 않을 경우 브라우저 닫기 전까지 유지된다        
        if (maxAge != 0) {
            cookie.setMaxAge(maxAge);
        }

        // 쿠키 저장
        response.addCookie(cookie);
    }

    /**
     * 쿠키값을 저장한다. path값을 "/"로 강제 설정
     * @param name - 쿠키 이름
     * @param value - 쿠키 값
     * @param maxAge - 쿠키 유효 시간 (0이면 지정 안함, 음수일 경우 즉시 삭제)
     * @param domain - 쿠키 도메인
     * 
     * @see #writeCookie(String, String, int, String, String)
     */
    public void writeCookie(String name, String value, int maxAge, String domain) throws Exception {
        this.writeCookie(name, value, maxAge, domain, "/");
    }

    /**
     * 쿠키값을 저장한다. path값을 "/"로, domain을 null로 강제 설정
     * @param name - 쿠키 이름
     * @param value - 쿠키 값
     * @param maxAge - 쿠키 유효 시간 (0이면 지정 안함, 음수일 경우 즉시 삭제)
     * 
     * @see #writeCookie(String, String, int, String, String)
     */
    public void writeCookie(String name, String value, int maxAge) throws Exception {
        this.writeCookie(name, value, maxAge, null, "/");
    }

    /**
     * 쿠키값을 저장한다. path값을 "/"로, domain을 null로, maxAge를 0으로 강제 설정
     * @param name - 쿠키 이름
     * @param value - 쿠키 값
     * @param maxAge - 쿠키 유효 시간 (0이면 지정 안함, 음수일 경우 즉시 삭제)
     * 
     * @see #writeCookie(String, String, int, String, String)
     */
    public void writeCookie(String name, String value) throws Exception {
        this.writeCookie(name, value, 0, null, "/");
    }

    /** 
     * 쿠키값을 삭제한다
     * @param name - 쿠키 이름
     */
    public void deleteCookie(String name) throws Exception {
        this.writeCookie(name, null, -1, null, "/");
    }


    /**
     * HTTP 상태코드를 설정하고 메시지를 출력한 후, 지정된 페이지로 이동한다
     * 이동할 페이지가 없을 경우 이전 페이지로 이동한다
     * @param statusCode - HTTP 상태코드
     * @param url - 이동할 URL
     * @param message - 출력할 메시지
     */
    public void redirect(int statusCode, String url, String message) {
        response.setStatus(statusCode);
        response.setContentType("text/html; charset=UTF-8");

        PrintWriter out = null;
        try {
            out = response.getWriter();
        } catch (IOException e) {
            log.error("응답 출력 스트림 생성 실패", e);
            return;
        }
        
        // 메시지가 있으면 메시지내용 alert
        if (message!=null && !message.equals("")) {
            out.print("<script>");
            out.print("alert('" + message + "');");
            out.print("</script>");
        }

        // url 있으면 url로 강제이동
        if (url!=null && !url.equals("")) {
            out.println("<meta http-equiv='refresh' content='0; url=" + url + "' />");
        } else {
            out.println("<script>");
            out.println("history.back()");
            out.println("</script>");
        }

        out.flush();
    }

    /**
     * HTTP 상태코드를 200으로 설정하고 메시지를 출력한 후, 지정된 페이지로 이동한다
     * @param url - 이동할 URL
     * @param message - 출력할 메시지
     */
    public void redirect (String url, String message) {
        this.redirect(200, url, message);
    }

    /**
     * HTTP 상태코드를 200으로 설정하고 메시지를 출력 없이, 지정된 페이지로 이동한다
     * @param url - 이동할 URL
     */
    public void redirect (String url) {
        this.redirect(200, url, null);
    }

    /**
     * HTTP 상태코드를 설정하고 메시지를 출력 없이, 지정된 페이지로 이동한다
     * @param statusCode - HTTP 상태코드
     * @param url - 이동할 URL
     */
    public void redirect (int statusCode, String url) {
        this.redirect(statusCode, url, null);
    }

    /**
     * 파라미터가 잘못된 경우에 호출할 이전 페이지 이동 기능
     * @param e - 에러정보를 담고 있는 객체. 
     *            Exception으로 선언했으므로 어떤 하위 객체가 전달되더라도 형변환되어 받는다
     */
    public void badRequest(Exception e) {
        this.redirect(403, null, e.getMessage());
    }

    /**
     * 파라미터가 잘못된 경우에 호출할 이전 페이지 이동 기능
     * @param message - 개발자가 직접 전달하는 에러 메시지
     */
    public void badRequest(String message) {
        this.redirect(403, null, message);
    }

    /** 
     * JAVA혹은 SQL쪽에서 잘못된 경우에 호출할 이전 페이지 이동 기능
     * @param e - 에러정보를 담고 있는 객체. 
     *            Exception으로 선언했으므로 어떤 하위 객체가 전달되더라도 형변환되어 받는다
     */
    public void serverError(Exception e) {
        String message = e.getMessage().trim().replace("'", "\\'").split(System.lineSeparator())[0];
        this.redirect(500, null, message);
    }
    
    /**
     * JAVA 혹은 SQL쪽에서 잘못된 경우에 호출할 이전 페이지 이동 기능
     * @param message - 개발자가 직접 전달하는 에러 메시지
     */
    public void serverError(String message) {
        this.redirect(500, null, message);
    }
}

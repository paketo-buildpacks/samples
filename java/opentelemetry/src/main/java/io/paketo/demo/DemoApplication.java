package io.paketo.demo;

import java.util.List;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.web.reactive.function.server.RouterFunction;
import org.springframework.web.reactive.function.server.RouterFunctions;
import org.springframework.web.reactive.function.server.ServerResponse;

@SpringBootApplication
public class DemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}

	@Bean
	RouterFunction<ServerResponse> routerFunction() {
		return RouterFunctions.route()
			.GET("/", request -> ServerResponse.ok().bodyValue("Hello, OpenTelemetry!"))
			.GET("/config", request -> ServerResponse.ok().bodyValue(List.of(
				"OTEL_JAVAAGENT_ENABLED=" + System.getenv("OTEL_JAVAAGENT_ENABLED"),
				"OTEL_SERVICE_NAME=" + System.getenv("OTEL_SERVICE_NAME")
			)))
			.build();
	}

}

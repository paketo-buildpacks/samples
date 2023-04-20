package io.paketo.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Flux;

import java.util.Arrays;

@SpringBootApplication
public class DemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}

	@RestController
	public static class StaticController {

		@GetMapping("/products")
		public Flux<Product> getAllProducts() {
			return Flux.fromIterable(Arrays.asList(
					new Product(1L, "Shovel", 10),
					new Product(2L, "Winter Coat", 100),
					new Product(3L, "Gloves", 20),
					new Product(4L, "Montréal Poutine", 20)
			));
		}

	}
}

record Product(Long id, String name, Integer price){}
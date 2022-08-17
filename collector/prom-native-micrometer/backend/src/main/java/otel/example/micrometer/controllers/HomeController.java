package otel.example.micrometer.controllers;

import otel.example.micrometer.entity.Greeting;
import otel.example.micrometer.repository.GreetingRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class HomeController {

    @Autowired
    private GreetingRepository repository;

    @GetMapping("/")
    public String showHome(String name, Model model) {
        Greeting dockerGreeting = repository.findById(1).orElse(new Greeting("Not Found 😕"));
        model = model.addAttribute("name", dockerGreeting.getName());
        return "home";
    }

}

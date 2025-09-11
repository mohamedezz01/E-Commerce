package entity;

import jakarta.persistence.*;
import org.hibernate.annotations.UpdateTimestamp;

import java.time.Instant;

@Entity
@Table(name = "inventory")
public class Inventory {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne
    @JoinColumn(name = "product_id",nullable = false,unique = true)
    private Product product;

    private Integer available;
    private Integer reserved;
    @UpdateTimestamp
    private Instant updatedAt;

}

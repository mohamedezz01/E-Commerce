package entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.CreationTimestamp;

import java.io.Serializable;
import java.math.BigDecimal;
import java.time.Instant;

@Entity
@Table(name = "coupons")
public class Coupon {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(unique = true, nullable = false)
    private String code;

    private String description;

    @Column(nullable = false)
    private String discountType; // percentage or fixed

    private BigDecimal discountValue;
    private Integer maxUses;
    private Integer usedCount = 0;
    private Instant expiresAt;
    private Boolean active = true;

    @CreationTimestamp
    private Instant createdAt;
}


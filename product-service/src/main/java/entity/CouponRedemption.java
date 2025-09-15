package entity;

import jakarta.persistence.*;
import org.hibernate.annotations.CreationTimestamp;

import java.time.Instant;

@Entity
@Table(name = "coupon_redemptions")
@IdClass(CouponRedemptionId.class)
public class CouponRedemption {
    @Id
    @ManyToOne @JoinColumn(name = "coupon_id")
    private Coupon coupon;

    @Id
    @ManyToOne @JoinColumn(name = "order_id")
    private Order order;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;

    @CreationTimestamp
    private Instant redeemedAt;
}

package entity;

import jakarta.persistence.*;
import org.hibernate.annotations.CreationTimestamp;

import java.time.Instant;

@Entity
@Table(name = "audit_logs")
public class AuditLog {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne @JoinColumn(name = "user_id")
    private User user;

    private String eventType;

    @Column(columnDefinition = "jsonb")
    private String eventData;

    private String ipAddress;

    @CreationTimestamp
    private Instant createdAt;
}

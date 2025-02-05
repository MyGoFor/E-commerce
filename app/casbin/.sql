INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES ('p', 'checkout', 'place_order', 'use');
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES ('p', 'checkout', 'mark_order_paid', 'use');
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES ('p', 'checkout', 'empty_cart', 'use');
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES ('p', 'checkout', 'charge', 'use');
INSERT INTO casbin_rule (ptype, v0, v1) VALUES ('g', 'checkout', 'checkout');

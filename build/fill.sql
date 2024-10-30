-- Вставляем данные
INSERT INTO users (login, password, is_admin) VALUES
('пользователь1', 'пароль1', false),
('админ', 'пароль_админ', true);

INSERT INTO task_items (image, title, minutes, description, answ) VALUES
('http://127.0.0.1:9000/prog/num1.png', 'Номер 3828 Демидович', 10, 'Определить аналитическое выражение для данного интеграла, применяя методы теории интегралов.', 'http://127.0.0.1:9000/prog/3828answ.png'),
('http://127.0.0.1:9000/prog/num2.png', 'Номер 3805 Демидович', 4, 'Найти значение этого интеграла, используя знания о нормальном распределении и преобразованиях Гаусса.', 'http://127.0.0.1:9000/prog/3805answ.png'),
('http://127.0.0.1:9000/prog/num3.png', 'Номер 3801 Демидович', 14, 'Решить данный определённый интеграл, используя подходящие методики интегрирования, такие как тригонометрические подстановки или интегрирование по частям.', 'http://127.0.0.1:9000/prog/3801answ.png'),
('http://127.0.0.1:9000/prog/num1.png', 'Номер 3801 Демидович', 79, 'Решить данный определённый интеграл, используя подходящие методики интегрирования, такие как тригонометрические подстановки или интегрирование по частям.', 'http://127.0.0.1:9000/prog/3801answ.png'),
('http://127.0.0.1:9000/prog/num1.png', 'Номер 3801 Демидович', 20, 'Решить данный определённый интеграл, используя подходящие методики интегрирования, такие как тригонометрические подстановки или интегрирование по частям.', 'http://127.0.0.1:9000/prog/3801answ.png');

-- Создание функции для проверки черновиковitem_r
CREATE FUNCTION check_draft_request() RETURNS trigger AS $$
BEGIN
    IF (SELECT COUNT(*) FROM lesson_requests WHERE user_id = NEW.user_id AND status = 'черновик') > 0 THEN
        RAISE EXCEPTION 'User already has a draft request';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание триггера для проверки статуса черновиков
CREATE TRIGGER draft_request_trigger
    BEFORE INSERT OR UPDATE ON lesson_requests
    FOR EACH ROW
    WHEN (NEW.status = 'черновик')
EXECUTE FUNCTION check_draft_request();
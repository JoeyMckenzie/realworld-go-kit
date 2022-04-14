-- seed users
insert into users (created_at, updated_at, username, email, password, bio, image)
values (current_timestamp, current_timestamp, 'user1', 'user1@gmail.com', 'user1', 'user1 bio', 'user1 image'),
       (current_timestamp, current_timestamp, 'user2', 'user2@gmail.com', 'user2', 'user2 bio', 'user2 image'),
       (current_timestamp, current_timestamp, 'user3', 'user3@gmail.com', 'user3', 'user3 bio', 'user3 image');

-- seed articles
insert into articles (created_at, updated_at, title, slug, description, body, user_id)
values (current_timestamp, current_timestamp, 'user1 title1', 'user1-slug1', 'user1 description1', 'user1 body1',
        (select id from users where username = 'user1')),
       (current_timestamp, current_timestamp, 'user1 title2', 'user2-slug2', 'user1-description2', 'user1-body2',
        (select id from users where username = 'user1')),
       (current_timestamp, current_timestamp, 'user2 title', 'user2-slug', 'user2 description', 'user2 body',
        (select id from users where username = 'user2'));

-- seed tags
insert into tags (created_at, tag)
values (current_timestamp, 'tag1'),
       (current_timestamp, 'tag2'),
       (current_timestamp, 'tag3');

-- seed article tags
insert into article_tags (created_at, tag_id, article_id)
values (current_timestamp, (select id from tags where tag = 'tag1'), (select id from articles limit 1));

-- select *
-- from jianying_shot_template
-- where is_topic = true
--   and search_term_text like "%term%"
-- order by total_use_cnt desc limit 20
-- offset 3
-- ;


-- select template_id
--         ,max(topic_id) topic_id
-- from jianying_shot_template
-- where is_topic = true
--   and search_term_text like "%term%"
--   group by template_id
-- ;


-- select  template_id
--         ,max(top_id) top_id
-- from jianying_shot_template
-- where is_topic = true
--   and search_term_text like "%term%"
--   group by template_id
-- ;


-- select template_id
--     ,top_id
-- from jianying_shot_template
-- where is_topic = true
-- ;


select template_id
from jianying_shot_template
where is_topic = 1
and template_id = 2
;
